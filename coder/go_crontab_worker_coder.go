/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-04-30 07:03:28
# File Name: go_crontab_worker_coder.go
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

type GoCrontabWorkerCoder struct {
	opts []*Option
	tpls []*Tpl
}

func NewGoCrontabWorkerCoder() *GoCrontabWorkerCoder {
	o := &GoCrontabWorkerCoder{}
	o.opts = []*Option{
		&Option{
			Name:  "author",
			Usage: "Input author name:"},
		&Option{
			Name:  "project_name",
			Usage: "Input project name:"},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_crontab_worker/conf/app.toml"},
		&Tpl{Src: "templates/go_crontab_worker/conf/log.xml"},
		&Tpl{Src: "templates/go_crontab_worker/handlers/handlers.go"},
		&Tpl{Src: "templates/go_crontab_worker/handlers/default_handler.go"},
		&Tpl{Src: "templates/go_crontab_worker/libs/config/config.go"},
		&Tpl{Src: "templates/go_crontab_worker/libs/types.go"},
		&Tpl{Src: "templates/go_crontab_worker/models/models.go"},
		&Tpl{Src: "templates/go_crontab_worker/.gitignore"},
		&Tpl{Src: "templates/go_crontab_worker/control.sh"},
		&Tpl{Src: "templates/go_crontab_worker/main.go"},
	}
	return o
}

func (this *GoCrontabWorkerCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoCrontabWorkerCoder) Init() (err error) {
	return
}

func (this *GoCrontabWorkerCoder) Generate() (err error) {
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
		tpl.Dst = fmt.Sprintf("%s/%s", dir, MacroReplace(strings.TrimPrefix(tpl.Src, "templates/go_crontab_worker/"), fileNameMacros))

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

func (this *GoCrontabWorkerCoder) getMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__AUTHOR__", Val: GetOptionValueByKey(this.opts, "author")},
		&Macro{Key: "__CREATE_DATETIME__", Val: time.Now().Format("2006-01-02 15:04:05")},
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
