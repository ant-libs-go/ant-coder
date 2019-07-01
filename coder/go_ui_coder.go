/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-04-30 02:37:16
# File Name: go_ui_coder.go
# Description:
####################################################################### */

package coder

import (
	"time"

	"github.com/ant-libs-go/util"
)

type GoUiCoder struct {
	opts []*Option
	tpls []*Tpl
}

func NewGoUiCoder() *GoUiCoder {
	o := &GoUiCoder{}
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
		&Tpl{Src: "templates/go_ui/conf/app.toml", Dst: "conf/app.toml"},
		&Tpl{Src: "templates/go_ui/conf/log.xml", Dst: "conf/log.xml"},
		&Tpl{Src: "templates/go_ui/controllers/controllers.go", Dst: "controllers/controllers.go"},
		&Tpl{Src: "templates/go_ui/controllers/site_controller.go", Dst: "controllers/site_controller.go"},
		&Tpl{Src: "templates/go_ui/libs/config/config.go", Dst: "libs/config/config.go"},
		&Tpl{Src: "templates/go_ui/libs/middlewares/logger.go", Dst: "libs/middlewares/logger.go"},
		&Tpl{Src: "templates/go_ui/libs/context.go", Dst: "libs/context.go"},
		&Tpl{Src: "templates/go_ui/libs/enums.go", Dst: "libs/enums.go"},
		&Tpl{Src: "templates/go_ui/libs/types.go", Dst: "libs/types.go"},
		&Tpl{Src: "templates/go_ui/models/models.go", Dst: "models/models.go"},
		&Tpl{Src: "templates/go_ui/.gitignore", Dst: ".gitignore"},
		&Tpl{Src: "templates/go_ui/control.sh", Dst: "control.sh"},
		&Tpl{Src: "templates/go_ui/main.go", Dst: "main.go"},
		&Tpl{Src: "templates/go_ui/server.go", Dst: "server.go"},
	}
	return o
}

func (this *GoUiCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoUiCoder) GetTpls() (r []*Tpl) {
	return this.tpls
}

func (this *GoUiCoder) Init() (err error) {
	return
}

func (this *GoUiCoder) GetBaseDirName() (r string) {
	return GetOptionValueByKey(this.opts, "project_name")
}

func (this *GoUiCoder) GetMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
