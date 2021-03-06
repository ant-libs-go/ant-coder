/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-11-14 12:50:43
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ant-coder/coder"
	"ant-coder/coder/config"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var coders = map[string]coder.Coder{
	"go_model":          coder.NewGoModelCoder(),
	"go_ui":             coder.NewGoUiCoder(),
	"go_loop_worker":    coder.NewGoLoopWorkerCoder(),
	"go_crontab_worker": coder.NewGoCrontabWorkerCoder(),
	"go_rpcx_server":    coder.NewGoRpcxServerCoder(),
}
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	verbose     = flag.String("v", "false", "enable verbose logging [false]")
	scene       string
)

func init() {
	var scenes []string
	for scene, _ := range coders {
		scenes = append(scenes, scene)
	}
	flag.StringVar(&scene, "s", "", fmt.Sprintf("coder scene (options: %s)", strings.Join(scenes, "|")))
	flag.Parse()

	if len(*pwd) == 0 {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", __version__)
	os.Setenv("WORKDIR", *pwd)
	os.Setenv("VERBOSE", *verbose)

	if len(scene) == 0 {
		fmt.Println("you must specify `-s` option")
		os.Exit(-1)
	}
	if err := config.SetPathAndLoad(os.Getenv("HOME")); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	c, ok := coders[scene]
	if !ok {
		fmt.Println("you specify coder sense not support.")
		os.Exit(-1)
	}
	if err := coder.NewExecutor(c).Do(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
