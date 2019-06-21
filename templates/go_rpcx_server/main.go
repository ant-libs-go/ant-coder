/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"__PROJECT_NAME__/handlers"
	"__PROJECT_NAME__/libs/config"
	"__PROJECT_NAME__/models"

	"github.com/cihub/seelog"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	cfg         = flag.String("c", "conf/app.toml", "config file, relative path")
	srv         *server.Server
)

func init() {
	flag.Parse()

	if *pwd == "" {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", __version__)
	os.Setenv("WORKDIR", *pwd)
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		fmt.Printf("Signal %d received, App is about to stop...\n", sig)
		srv.Close()
		fmt.Printf("App has gone away\n")
		//time.Sleep(time.Second)
		//os.Exit(0)
	}()
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()
	// configuration
	fmt.Printf("Using configuration %s\n", *cfg)
	if strings.HasPrefix(*cfg, "/") == false {
		*cfg = path.Join(os.Getenv("WORKDIR"), *cfg)
	}
	if err := config.SetFileAndLoad(*cfg); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// logger
	logFile := config.Get().Basic.LogFile
	fmt.Printf("Using log configuration %s\n", logFile)
	if strings.HasPrefix(logFile, "/") == false {
		logFile = path.Join(os.Getenv("WORKDIR"), logFile)
	}
	var err error
	var logger seelog.LoggerInterface
	if logger, err = seelog.LoggerFromConfigAsFile(logFile); err != nil {
		fmt.Printf("Log configuration parse error: %s\n", err)
		os.Exit(-1)
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	// GC
	debug.SetGCPercent(config.Get().Gc.Percent)

	// init models
	models.Init()

	// register stop listener
	registerSignalHandler()

	// rpc server
	srv = server.NewServer()
	go func() {
		if err := srv.Serve("tcp", config.Get().Basic.Port); err != nil {
			fmt.Printf("App start error: %s\n", err)
			os.Exit(-1)
		}
	}()
	time.Sleep(time.Second)
	fmt.Printf("Listening on %s\n", srv.Address().String())

	host, err1 := os.Hostname()
	_, port, err2 := net.SplitHostPort(srv.Address().String())
	if err1 != nil || err2 != nil {
		fmt.Printf("Host or Port parse error: %s, %s\n", err1, err2)
		os.Exit(-1)
	}
	register := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf("tcp@%s:%s", host, port),
		ConsulServers:  []string{"127.0.0.1:8500"},
		BasePath:       config.Get().Basic.Node,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	srv.Plugins.Add(register)

	if err := srv.RegisterName("Server", handlers.NewServiceImpl(), ""); err != nil {
		fmt.Printf("Build service error: %s\n", err)
		os.Exit(-1)
	}

	if err := register.Start(); err != nil {
		fmt.Printf("Register to registry error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("App start ok, running as %d\n", os.Getpid())
	select {}
}
