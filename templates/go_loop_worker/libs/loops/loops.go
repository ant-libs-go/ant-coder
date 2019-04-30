/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-12-20 11:26:21
# File Name: loops.go
# Description:
####################################################################### */

package loops

import (
	"os"
	"time"
)

type Entry struct {
	Spec time.Duration
	Job  func()
}

type Loops struct {
	ExeDir  string
	entries []*Entry
	running bool
}

func New() *Loops {
	o := &Loops{
		ExeDir:  os.Getenv("EXECDIR"),
		running: true,
	}
	return o
}

func (this *Loops) AddFunc(spec time.Duration, cmd func()) {
	entry := &Entry{
		Spec: spec,
		Job:  cmd,
	}
	this.entries = append(this.entries, entry)
}

func (this *Loops) Start() {
	for _, entry := range this.entries {
		go func(entry *Entry) {
			for this.running {
				entry.Job()
				time.Sleep(entry.Spec)
			}
		}(entry)
	}
}

func (this *Loops) Stop() {
	this.running = false
}
