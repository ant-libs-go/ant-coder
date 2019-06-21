/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: default_handler.go
# Description:
####################################################################### */

package handlers

import (
	"github.com/ant-libs-go/util/logs"
	"gitlab.com/feichi/fcad_thrift/libs/go/common"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
	services "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_services"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_types"
)

type DefaultServiceImpl struct {
}

func NewDefaultServiceImpl() *DefaultServiceImpl {
	o := &DefaultServiceImpl{}
	return o
}

func (this *DefaultServiceImpl) GetByIds(req *services.GetMediaByIdsRequest, log *logs.SessLog) (r *services.GetMediaByIdsResponse) {
	r = &services.GetByIdsResponse{
		Header: &common.Header{Code: enums.ResponseCode_OK},
		Body:   map[int32]*types.Media{}}

	/*
		_, ms, err := query.New().GetByIds(req.Body)
		if err != nil {
			log.Warnf("server exception, %v", err)
			r.Header.Code = enums.ResponseCode_SERVER_ERROR
			return
		}

		for _, m := range ms {
			r.Body[m.Id] = m.Thrift()
		}
	*/
	return
}
