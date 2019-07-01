/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2019-02-20 14:27:22
# File Name: coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
	"os"
	"path"
)

type Option struct {
	Name    string
	Default string
	Value   string
	Usage   string
	Cache   bool
}

type Tpl struct {
	Src string
	Dst string
}
type Macro struct {
	Key string
	Val string
}

type Coder interface {
	GetOptions() (r []*Option)
	GetTpls() (r []*Tpl)
	Init() (err error)
	GetBaseDirName() (r string)
	GetMacros() (r []*Macro, r2 []*Macro, err error)
}

type Executor struct {
	coder Coder
	opts  []*Option
}

func NewExecutor(coder Coder) *Executor {
	o := &Executor{}
	o.coder = coder
	o.opts = o.coder.GetOptions()
	return o
}

func (this *Executor) Do() (err error) {
	if err = this.inputOptions(); err != nil {
		return fmt.Errorf("input options err: %+v", err)
	}
	fmt.Println("\nValid options...")
	if err = this.validOptions(); err != nil {
		return fmt.Errorf("\nvalid options err: %+v\n", err)
	}
	fmt.Println("\nInit coder...")
	if err = this.coder.Init(); err != nil {
		return fmt.Errorf("\ncoder init err: %+v\n", err)
	}
	fmt.Println("\nCoder Generate...")
	if err = this.generate(); err != nil {
		return fmt.Errorf("\ncoder generate err: %+v\n", err)
	}
	fmt.Println("\nDone...")
	return
}

func (this *Executor) inputOptions() (err error) {
	for _, opt := range this.opts {
		var tips string
		if len(opt.Default) > 0 {
			tips = fmt.Sprintf("%s [%s] ", opt.Usage, opt.Default)
		} else {
			tips = fmt.Sprintf("%s ", opt.Usage)
		}

		opt.Value = Scan(tips)
		if len(opt.Value) > 0 {
			continue
		}
		opt.Value = opt.Default
	}
	return
}

func (this *Executor) validOptions() (err error) {
	var opt *Option
	defer func() {
		if err == nil {
			return
		}
		fmt.Println(fmt.Sprintf("....... options#%s\t[no]", opt.Name))
	}()

	for _, opt = range this.opts {
		if len(opt.Value) == 0 {
			err = fmt.Errorf("%s can't be blank", opt.Name)
			return
		}
		fmt.Println(fmt.Sprintf("....... options#%s\t[ok]", opt.Name))
	}
	return
}

func (this *Executor) generate() (err error) {
	// mkdir base dir
	dir := fmt.Sprintf("%s/%s", os.Getenv("WORKDIR"), this.coder.GetBaseDirName())
	if err = Mkdir(dir); err != nil {
		fmt.Println(fmt.Sprintf("....... directory#%s mkdir\t[no]", dir))
		return
	}
	fmt.Println("....... directory mkdir\t[ok]")

	// render tpl
	fileNameMacros, fileContMacros, err := this.coder.GetMacros()
	if err != nil {
		return err
	}
	for _, tpl := range this.coder.GetTpls() {
		tpl.Dst = fmt.Sprintf("%s/%s", dir, MacroReplace(tpl.Dst, fileNameMacros))
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
