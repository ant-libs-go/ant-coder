/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: types.go
# Description:
####################################################################### */

package libs

type Params interface {
	GetCommonParams() *CommonParams
}

type CommonParams struct {
	Token string `form:"token" binding:""`
}

type Pager struct {
	Page     int `form:"page"`
	LastId   int `form:"last_id"`
	PageSize int `form:"page_size"`
}

func (this *CommonParams) GetToken() string {
	return this.Token
}

func (this *CommonParams) GetCommonParams() *CommonParams {
	return this
}

type LoginParams struct {
	CommonParams
	Username string `form:"username" binding:""`
	Passport string `form:"passport" binding:""`
}

func (this *LoginParams) GetCommonParams() *CommonParams {
	return &this.CommonParams
}

type LoginResult struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
