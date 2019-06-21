/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: handlers.go
# Description:
####################################################################### */

package handlers

import (
	"context"

	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/feichi/fcad_thrift/libs/go/common"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
	services "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_services"
)

type ServiceImpl struct {
	DefaultHandler *DefaultServiceImpl
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{
		DefaultHandler: NewDefaultServiceImpl(),
	}
}

func (this *ServiceImpl) before(header *common.Header) (log *logs.SessLog) {
	if len(header.Sessid) == 0 {
		header.Sessid = uuid.NewV4().String()
	}
	log = logs.New(header.Sessid)
	return
}

func (this *ServiceImpl) GetByIds(ctx context.Context, req *services.GetMediaByIdsRequest, resp *services.GetMediaByIdsResponse) (err error) {
	log := this.before(req.Header)
	log.Infof("Request type: GetByIds, req: %v", req)
	resp.Header = &common.Header{Sessid: req.Header.Sessid, Code: enums.ResponseCode_OK, Metadata: req.Header.Metadata}

	r := this.DefaultHandler.GetByIds(req, log)
	resp.Header.Code = r.Header.Code
	resp.Body = r.Body
	if resp.Header.Code != enums.ResponseCode_OK {
		log.Warnf("Do error, code: %v", resp.Header.Code)
		return
	}

	log.Infof("Run success: %v", resp)
	return
}
