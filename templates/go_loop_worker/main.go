/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	"__PROJECT_NAME__/handlers"
	"__PROJECT_NAME__/libs/config"
	"__PROJECT_NAME__/libs/loops"
	"__PROJECT_NAME__/models"

	"github.com/cihub/seelog"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	cfg         = flag.String("c", "conf/app.toml", "config file, relative path")
	wg          sync.WaitGroup
	loopws      *loops.Loops
	exit        = make(chan int)
)

func init() {
	flag.Parse()

	if *pwd == "" {
		*pwd, _ = os.Getwd()
	}
	os.Setenv("VERSION", __version__)
	os.Setenv("WORKDIR", *pwd)
}

func registerSignalHandler() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		fmt.Printf("Signal %d received, App is about to stop...\n", sig)
		go loopws.Stop()
		wg.Wait()
		time.Sleep(time.Second)
		fmt.Printf("App has gone away\n")
		close(exit)
	}()
}

func main() {
	fmt.Printf("Running in %s\n", *pwd)

	// configuration
	fmt.Printf("Using configuration %s\n", *cfg)
	if strings.HasPrefix(*cfg, "/") == false {
		*cfg = path.Join(os.Getenv("WORKDIR"), *cfg)
	}
	if err := config.SetFileAndLoad(*cfg); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// logger
	logFile := config.Get().Basic.LogFile
	fmt.Printf("Using log configuration %s\n", logFile)
	if strings.HasPrefix(logFile, "/") == false {
		logFile = path.Join(os.Getenv("WORKDIR"), logFile)
	}
	var err error
	var logger seelog.LoggerInterface
	if logger, err = seelog.LoggerFromConfigAsFile(logFile); err != nil {
		fmt.Printf("Log configuration parse error: %s\n", err)
		os.Exit(-1)
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	// GC
	debug.SetGCPercent(config.Get().Gc.Percent)

	// init models
	models.Init()

	// register stop listener
	registerSignalHandler()

	loopws = loops.New()
	loopws.AddFunc(time.Second, func() { wg.Add(1); defer wg.Done(); handlers.NewMarkOverdueHandler().Run() })
	// add handler
	loopws.Start()

	fmt.Printf("App start ok, running as %d\n", os.Getpid())
	select {
	case <-exit:
	}
}
