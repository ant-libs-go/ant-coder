/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: types.go
# Description:
####################################################################### */

package libs

type ResponseCode int64

const (
	ResponseCode_OK           ResponseCode = 200
	ResponseCode_BAD_REQUEST  ResponseCode = 400
	ResponseCode_SERVER_ERROR ResponseCode = 500
	ResponseCode_TIMEOUT      ResponseCode = 504
)

type Header struct {
	Requester string
	Sessid    string
	Timestamp int64
	Version   int32
	Operator  int64
	Code      ResponseCode
	Metadata  map[string]string
}

type GetByIdsRequest struct {
	Header *Header
	Body   []int32
}

type GetByIdsResponse struct {
	Header *Header
	Body   map[int32]string
}
