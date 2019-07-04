/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-06-21 03:57:08
# File Name: go_rpcx_server_coder.go
# Description:
####################################################################### */

package coder

import (
	"time"

	"github.com/ant-libs-go/util"
)

type GoRpcxServerCoder struct {
	name string
	opts []*Option
	tpls []*Tpl
}

func NewGoRpcxServerCoder() *GoRpcxServerCoder {
	o := &GoRpcxServerCoder{}
	o.name = "go_rpcx_server"
	o.opts = []*Option{
		&Option{
			Name:  "author",
			Usage: "Input author name:",
			Cache: true},
		&Option{
			Name:  "project_name",
			Usage: "Input project name:"},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_rpcx_server/conf/app.toml", Dst: "conf/app.toml"},
		&Tpl{Src: "templates/go_rpcx_server/conf/log.xml", Dst: "conf/log.xml"},
		&Tpl{Src: "templates/go_rpcx_server/handlers/handlers.go", Dst: "handlers/handlers.go"},
		&Tpl{Src: "templates/go_rpcx_server/handlers/handlers_test.go", Dst: "handlers/handlers_test.go"},
		&Tpl{Src: "templates/go_rpcx_server/handlers/default_handler.go", Dst: "handlers/default_handler.go"},
		&Tpl{Src: "templates/go_rpcx_server/libs/config/config.go", Dst: "libs/config/config.go"},
		&Tpl{Src: "templates/go_rpcx_server/libs/types.go", Dst: "libs/types.go"},
		&Tpl{Src: "templates/go_rpcx_server/models/models.go", Dst: "models/models.go"},
		&Tpl{Src: "templates/go_rpcx_server/.gitignore", Dst: ".gitignore"},
		&Tpl{Src: "templates/go_rpcx_server/control.sh", Dst: "control.sh"},
		&Tpl{Src: "templates/go_rpcx_server/main.go", Dst: "main.go"},
	}
	return o
}

func (this *GoRpcxServerCoder) GetName() (r string) {
	return this.name
}

func (this *GoRpcxServerCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoRpcxServerCoder) GetTpls() (r []*Tpl) {
	return this.tpls
}

func (this *GoRpcxServerCoder) Init() (err error) {
	return
}

func (this *GoRpcxServerCoder) GetBaseDirName() (r string) {
	return GetOptionValueByKey(this.opts, "project_name")
}

func (this *GoRpcxServerCoder) GetMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
