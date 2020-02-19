/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: types.go
# Description:
####################################################################### */

package libs

type Params interface {
	GetUserId() int32
}

type CommonParams struct {
	UserId int32 `form:"user_id" binding:""`
}

func (this *CommonParams) GetUserId() int32 {
	return this.UserId
}

type LoginReq struct {
	CommonParams
	Username string `form:"username" binding:""`
	Passport string `form:"passport" binding:""`
}

type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
