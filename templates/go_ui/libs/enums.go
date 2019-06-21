/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
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
