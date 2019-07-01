/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-02-20 18:12:02
# File Name: misc.go
# Description:
####################################################################### */

package coder

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ant-coder/templates"

	"github.com/ant-libs-go/util"
)

func GetOptionByKey(opts []*Option, key string) *Option {
	for _, opt := range opts {
		if opt.Name == key {
			return opt
		}
	}
	return nil
}

func GetOptionValueByKey(opts []*Option, key string) string {
	opt := GetOptionByKey(opts, key)
	if opt == nil {
		return ""
	}
	return opt.Value
}

func Scan(tips string) (r string) {
	inp := bufio.NewScanner(os.Stdin)
	fmt.Print(tips)
	inp.Scan()
	r = strings.TrimSpace(inp.Text())
	return
}

func Mkdir(dir string) (err error) {
	exist, isdir, err := util.PathExists(dir)
	if err != nil {
		return fmt.Errorf("check directory#%s is exist, %+v", dir, err)
	}
	if exist == true && isdir == false {
		return fmt.Errorf("check directory#%s is exist, already exist but not is directory", dir)
	}
	if exist == false {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("mkdir directory#%s, %+v", dir, err)
		}
	}
	return
}

func RenderTpl(tpl *Tpl, macros []*Macro) (err error) {
	exist, _, err := util.PathExists(tpl.Dst)
	if err != nil {
		return fmt.Errorf("render tpl file#%s, %+v", tpl.Dst, err)
	}
	if exist {
		if Scan(fmt.Sprintf("file#%s already exist, overwrite? [n] ", tpl.Dst)) != "y" {
			return
		}
	}
	b, err := templates.Asset(tpl.Src)
	if err != nil {
		return fmt.Errorf("render tpl file#%s, %+v", tpl.Dst, err)
	}
	err = util.WriteFile(MacroReplace(string(b), macros), tpl.Dst)
	if err != nil {
		return fmt.Errorf("render tpl file#%s, %+v", tpl.Dst, err)
	}
	return
}

func MacroReplace(str string, macros []*Macro) (r string) {
	r = str
	for _, macro := range macros {
		r = strings.Replace(r, macro.Key, macro.Val, -1)
	}
	return
}
