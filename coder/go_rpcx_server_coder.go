/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-06-21 03:57:08
# File Name: go_rpcx_server_coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ant-libs-go/util"
)

type GoRpcxServerCoder struct {
	opts []*Option
	tpls []*Tpl
}

func NewGoRpcxServerCoder() *GoRpcxServerCoder {
	o := &GoRpcxServerCoder{}
	o.opts = []*Option{
		&Option{
			Name:  "author",
			Usage: "Input author name:"},
		&Option{
			Name:  "project_name",
			Usage: "Input project name:"},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_rpcx_server/conf/app.toml"},
		&Tpl{Src: "templates/go_rpcx_server/conf/log.xml"},
		&Tpl{Src: "templates/go_rpcx_server/handlers/handlers.go"},
		&Tpl{Src: "templates/go_rpcx_server/handlers/default_handler.go"},
		&Tpl{Src: "templates/go_rpcx_server/libs/config/config.go"},
		&Tpl{Src: "templates/go_rpcx_server/libs/types.go"},
		&Tpl{Src: "templates/go_rpcx_server/models/models.go"},
		&Tpl{Src: "templates/go_rpcx_server/.gitignore"},
		&Tpl{Src: "templates/go_rpcx_server/control.sh"},
		&Tpl{Src: "templates/go_rpcx_server/main.go"},
		&Tpl{Src: "templates/go_rpcx_server/test/client.go"},
	}
	return o
}

func (this *GoRpcxServerCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoRpcxServerCoder) Init() (err error) {
	return
}

func (this *GoRpcxServerCoder) Generate() (err error) {
	// mkdir dir
	dir := fmt.Sprintf("%s/%s", os.Getenv("WORKDIR"), GetOptionValueByKey(this.opts, "project_name"))
	if err = Mkdir(dir); err != nil {
		fmt.Println(fmt.Sprintf("....... directory#%s mkdir\t[no]", dir))
		return
	}
	fmt.Println("....... directory mkdir\t[ok]")

	// render tpl
	fileNameMacros, fileContMacros, err := this.getMacros()
	if err != nil {
		return err
	}
	for _, tpl := range this.tpls {
		tpl.Dst = fmt.Sprintf("%s/%s", dir, MacroReplace(strings.TrimPrefix(tpl.Src, "templates/go_rpcx_server/"), fileNameMacros))

		if err = Mkdir(path.Dir(tpl.Dst)); err != nil {
			fmt.Println(fmt.Sprintf("....... directory#%s mkdir\t[no]", path.Dir(tpl.Dst)))
			return
		}

		if err = RenderTpl(tpl, fileContMacros); err != nil {
			return
		}
	}
	fmt.Println("....... render template\t[ok]")

	return
}

func (this *GoRpcxServerCoder) getMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
