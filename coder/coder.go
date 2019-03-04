/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2019-02-20 14:27:22
# File Name: coder.go
# Description:
####################################################################### */

package coder

import (
	"fmt"
)

type Option struct {
	Name    string
	Default string
	Value   string
	Usage   string
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
	Init() (err error)
	Generate() (err error)
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
	if err = this.coder.Generate(); err != nil {
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