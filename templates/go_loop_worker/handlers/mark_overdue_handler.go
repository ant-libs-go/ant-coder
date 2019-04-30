/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-12-25 17:35:54
# File Name: mark_overdue_handler.go
# Description:
####################################################################### */

package handlers

import (
	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
)

type MarkOverdueHandler struct {
	Name string
	log  *logs.SessLog
}

func NewMarkOverdueHandler() *MarkOverdueHandler {
	o := &MarkOverdueHandler{}
	o.Name = "MARK_OVERDUE_HANDLER"
	o.log = logs.New(uuid.NewV4().String())
	return o
}

func (this *MarkOverdueHandler) Run() {
	this.log.Infof("Recall#%s...", this.Name)

	// code in here...

	this.log.Infof("recall#%s finish", this.Name)
}
