/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2019-07-03 17:49:25
# File Name: handlers_test.go
# Description:
####################################################################### */

//go test -v -tags consul ./handlers -run TestGetByIds -pwd=`pwd` -registry=false
package handlers

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"__PROJECT_NAME__/libs/config"
	"__PROJECT_NAME__/models"

	"github.com/ant-libs-go/util"
	"github.com/cihub/seelog"
	"github.com/smallnest/rpcx/client"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	defaultPwd, _ = os.Getwd()
	defaultNode   = "/__PROJECT_NAME__"
)

var (
	cli      client.XClient
	handler  = NewServiceImpl()
	pwd      = flag.String("pwd", defaultPwd, "work directory")
	cfg      = flag.String("cfg", "conf/app.toml", "config file, relative path")
	log      = flag.Bool("log", false, "show log?")
	node     = flag.String("node", defaultNode, "consul node name")
	registry = flag.Bool("registry", false, "do you want use registry?")
)

func TestMain(m *testing.M) {
	flag.Parse()

	fmt.Printf("Using configuration: %s\n", *cfg)
	fmt.Printf("Using registry: %v\n", *registry)
	// configuration
	if strings.HasPrefix(*cfg, "/") == false {
		*cfg = path.Join(*pwd, *cfg)
	}
	if err := config.SetFileAndLoad(*cfg); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// logger
	seelog.ReplaceLogger(seelog.Disabled)
	if *log == true {
		logFile := config.Get().Basic.LogFile
		fmt.Printf("Using log configuration %s\n", logFile)
		if strings.HasPrefix(logFile, "/") == false {
			logFile = path.Join(*pwd, logFile)
		}
		var err error
		var logger seelog.LoggerInterface
		if logger, err = seelog.LoggerFromConfigAsFile(logFile); err != nil {
			fmt.Printf("Log configuration parse error: %s\n", err)
			os.Exit(-1)
		}
		seelog.ReplaceLogger(logger)
	}
	defer seelog.Flush()

	// init models
	models.Init()
	if *log == false {
		models.Orm.ShowSQL(false)
	}

	if *registry == true {
		d := client.NewConsulDiscovery(*node, "Server", []string{"127.0.0.1:8500"}, nil)
		cli = client.NewXClient("Server", client.Failover, client.RandomSelect, d, client.DefaultOption)
		defer cli.Close()
	}
	os.Exit(m.Run())
}

func buildCommonHeader() *libs.Header {
	return &libs.Header{
		Requester: "test-client",
		Timestamp: time.Now().Unix(),
		Version:   1,
		Operator:  1,
		Metadata:  map[string]string{}}
}

func Call(fn, req, resp interface{}) (err error) {
	if *registry == true {
		fname := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		re, _ := regexp.Compile(`\.(?P<fn>[\w]+)-fm$`)
		ret, _ := util.FindStringSubmatch(re, fname)
		err = cli.Call(context.Background(), ret["fn"], req, resp)
	} else {
		args := []reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(req), reflect.ValueOf(resp)}
		reflect.ValueOf(fn).Call(args)
	}

	header := reflect.ValueOf(resp).Elem().FieldByName("Header").Interface().(*libs.Header)
	if header.Code != libs.ResponseCode_OK {
		err = fmt.Errorf("response code is %s, not is ok", header.Code)
	}
	fmt.Print(fmt.Sprintf(". err: %+v -> ", err))
	return
}

func TestGetByIds(t *testing.T) {
	req := &libs.GetByIdsRequest{
		Header: buildCommonHeader(),
		Body:   []int32{108, 109},
	}
	resp := &libs.GetByIdsResponse{}

	Convey("TestGetByIds", t, func() {
		Convey("TestGetByIds should return nil", func() {
			So(Call(handler.GetByIds, req, resp), ShouldBeNil)
		})
	})
}
