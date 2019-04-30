/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-14 00:32:52
# File Name: enums.go
# Description:
####################################################################### */

package libs

type HttpStatus int

const (
	HttpStatusOk        HttpStatus = 200
	HttpStatusParamsErr HttpStatus = 400 // 参数验证错误
	HttpStatusAuthErr   HttpStatus = 401 // token 过期
	HttpStatusNotFound  HttpStatus = 404
	HttpStatusServerErr HttpStatus = 500 // 服务器内部错误
)

type Role string

const (
	RoleAny   Role = ""
	RoleGuest Role = "?"
	RoleAuth  Role = "@"
)
