/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-04-30 07:03:28
# File Name: go_crontab_worker_coder.go
# Description:
####################################################################### */

package coder

import (
	"time"

	"github.com/ant-libs-go/util"
)

type GoCrontabWorkerCoder struct {
	name string
	opts []*Option
	tpls []*Tpl
}

func NewGoCrontabWorkerCoder() *GoCrontabWorkerCoder {
	o := &GoCrontabWorkerCoder{}
	o.name = "go_crontab_worker"
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
		&Tpl{Src: "templates/go_crontab_worker/conf/app.toml", Dst: "conf/app.toml"},
		&Tpl{Src: "templates/go_crontab_worker/conf/log.xml", Dst: "conf/log.xml"},
		&Tpl{Src: "templates/go_crontab_worker/handlers/handlers.go", Dst: "handlers/handlers.go"},
		&Tpl{Src: "templates/go_crontab_worker/handlers/default_handler.go", Dst: "handlers/default_handler.go"},
		&Tpl{Src: "templates/go_crontab_worker/libs/config/config.go", Dst: "libs/config/config.go"},
		&Tpl{Src: "templates/go_crontab_worker/libs/types.go", Dst: "libs/types.go"},
		&Tpl{Src: "templates/go_crontab_worker/models/models.go", Dst: "models/models.go"},
		&Tpl{Src: "templates/go_crontab_worker/.gitignore", Dst: ".gitignore"},
		&Tpl{Src: "templates/go_crontab_worker/control.sh", Dst: "control.sh"},
		&Tpl{Src: "templates/go_crontab_worker/main.go", Dst: "main.go"},
	}
	return o
}

func (this *GoCrontabWorkerCoder) GetName() (r string) {
	return this.name
}

func (this *GoCrontabWorkerCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoCrontabWorkerCoder) GetTpls() (r []*Tpl) {
	return this.tpls
}

func (this *GoCrontabWorkerCoder) Init() (err error) {
	return
}

func (this *GoCrontabWorkerCoder) GetBaseDirName() (r string) {
	return GetOptionValueByKey(this.opts, "project_name")
}

func (this *GoCrontabWorkerCoder) GetMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
