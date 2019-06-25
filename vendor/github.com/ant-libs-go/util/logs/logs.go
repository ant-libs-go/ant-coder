/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-13 08:08:55
# File Name: logs.go
# Description:
####################################################################### */

/**
 * import (
 * 	"os"
 * 	"github.com/cihub/seelog"
 * )
 * logger, err := seelog.LoggerFromConfigAsFile("log.xml")
 * if  err != nil {
 *     os.Exit(-1)
 * }
 * seelog.ReplaceLogger(logger)
 * defer seelog.Flush()
 *
 * log := logs.New()
 * log.Infof("this is a %s", "log")
 */

package logs

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/ant-libs-go/util"
	"github.com/cihub/seelog"
)

var (
	lock    sync.RWMutex
	entries map[string]*SessLog
)

type SessLog struct {
	Sessid string
	Logger seelog.LoggerInterface
}

func init() {
	entries = make(map[string]*SessLog)
}

/**
 * 必须调用Close方法
 */
func New(sessid string) *SessLog {
	if len(sessid) == 0 {
		sessid = strconv.Itoa(util.Goid())
	}
	o := &SessLog{}
	o.Sessid = sessid
	o.Logger = seelog.Current
	lock.Lock()
	entries[sessid] = o
	lock.Unlock()
	return o
}

func Get(sessid string) *SessLog {
	lock.RLock()
	defer lock.RUnlock()
	if v, ok := entries[sessid]; ok {
		return v
	}
	return nil
}

func Close(sessid string) {
	if len(sessid) == 0 {
		sessid = strconv.Itoa(util.Goid())
	}
	lock.Lock()
	delete(entries, sessid)
	lock.Unlock()
}

func (this *SessLog) Tracef(f string, v ...interface{}) {
	this.Logger.Tracef(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}

func (this *SessLog) Debugf(f string, v ...interface{}) {
	this.Logger.Debugf(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}

func (this *SessLog) Infof(f string, v ...interface{}) {
	this.Logger.Infof(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}

func (this *SessLog) Warnf(f string, v ...interface{}) {
	this.Logger.Warnf(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}

func (this *SessLog) Errorf(f string, v ...interface{}) {
	this.Logger.Errorf(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}

func (this *SessLog) Criticalf(f string, v ...interface{}) {
	this.Logger.Criticalf(fmt.Sprintf("[sid:%s] %s", this.Sessid, f), v...)
}
