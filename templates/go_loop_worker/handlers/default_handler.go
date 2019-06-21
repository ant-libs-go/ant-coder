/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: default_handler.go
# Description:
####################################################################### */

package handlers

import (
	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
)

type DefaultHandler struct {
	Name string
	log  *logs.SessLog
}

func NewDefaultHandler() *DefaultHandler {
	o := &DefaultHandler{}
	o.Name = "DEFAULT_HANDLER"
	o.log = logs.New(uuid.NewV4().String())
	return o
}

func (this *DefaultHandler) Run() {
	this.log.Infof("Recall#%s...", this.Name)

	// code in here...

	this.log.Infof("recall#%s finish", this.Name)
}
