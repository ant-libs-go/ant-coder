/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-11 14:35:24
# File Name: main.go
# Description:
####################################################################### */

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"__PROJECT_NAME__/libs/config"
	"__PROJECT_NAME__/models"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

// pass through when build project, go build -ldflags "main.__version__ 1.2.1" app
var (
	__version__ string
	pwd         = flag.String("d", "", "work directory")
	cfg         = flag.String("c", "conf/app.toml", "config file, relative path")
	srv         *http.Server
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
		fmt.Printf("Signal %d received, UiServer is about to stop...\n", sig)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		srv.Shutdown(ctx)
		fmt.Printf("UiServer has gone away\n")
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

	// http server
	if config.Get().Basic.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	var handler *gin.Engine
	if handler, err = NewUiServer(); err != nil {
		fmt.Printf("UiServer build error: %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("UIServer listening on %s\n", config.Get().Basic.Port)
	srv = &http.Server{Addr: config.Get().Basic.Port, Handler: handler}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("UiServer stopped: %s\n", err)
		os.Exit(-1)
	}
}
