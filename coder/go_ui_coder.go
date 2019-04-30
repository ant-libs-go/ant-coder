/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-04-30 02:37:16
# File Name: go_ui_coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
	"os"
	"path"
	"strings"

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
			Name:  "project_name",
			Usage: "Input project name:"},
	}
	o.tpls = []*Tpl{
		&Tpl{Src: "templates/go_ui/conf/app.toml"},
		&Tpl{Src: "templates/go_ui/conf/log.xml"},
		&Tpl{Src: "templates/go_ui/controllers/controllers.go"},
		&Tpl{Src: "templates/go_ui/controllers/site_controller.go"},
		&Tpl{Src: "templates/go_ui/libs/config/config.go"},
		&Tpl{Src: "templates/go_ui/libs/middlewares/logger.go"},
		&Tpl{Src: "templates/go_ui/libs/context.go"},
		&Tpl{Src: "templates/go_ui/libs/enums.go"},
		&Tpl{Src: "templates/go_ui/libs/types.go"},
		&Tpl{Src: "templates/go_ui/models/models.go"},
		&Tpl{Src: "templates/go_ui/.gitignore"},
		&Tpl{Src: "templates/go_ui/control.sh"},
		&Tpl{Src: "templates/go_ui/main.go"},
		&Tpl{Src: "templates/go_ui/server.go"},
	}
	return o
}

func (this *GoUiCoder) GetOptions() (r []*Option) {
	return this.opts
}

func (this *GoUiCoder) Init() (err error) {
	return
}

func (this *GoUiCoder) Generate() (err error) {
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
		tpl.Dst = fmt.Sprintf("%s/%s", dir, MacroReplace(strings.TrimPrefix(tpl.Src, "templates/go_ui/"), fileNameMacros))

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

func (this *GoUiCoder) getMacros() (fileNameMacros []*Macro, fileContMacros []*Macro, err error) {
	fileNameMacros = []*Macro{}
	fileContMacros = []*Macro{
		&Macro{Key: "__PROJECT_NAME__", Val: GetOptionValueByKey(this.opts, "project_name")},
		&Macro{Key: "__PROJECT_NAME_CAMEL__", Val: util.CamelString(GetOptionValueByKey(this.opts, "project_name"))},
	}

	return
}
