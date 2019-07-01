/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: default_handler.go
# Description:
####################################################################### */

package handlers

import (
	"__PROJECT_NAME__/libs"

	"github.com/ant-libs-go/util/logs"
)

type DefaultServiceImpl struct {
}

func NewDefaultServiceImpl() *DefaultServiceImpl {
	o := &DefaultServiceImpl{}
	return o
}

func (this *DefaultServiceImpl) GetByIds(req *libs.GetByIdsRequest, log *logs.SessLog) (r *libs.GetByIdsResponse) {
	r = &libs.GetByIdsResponse{
		Header: &libs.Header{Code: libs.ResponseCode_OK},
		Body:   map[int32]string{}}

	/*
		_, ms, err := query.New().GetByIds(req.Body)
		if err != nil {
			log.Warnf("server exception, %v", err)
			r.Header.Code = libs.ResponseCode_SERVER_ERROR
			return
		}

		for _, m := range ms {
			r.Body[m.Id] = m.Format()
		}
	*/
	return
}
