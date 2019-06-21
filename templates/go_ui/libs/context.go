/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: context.go
# Description:
####################################################################### */

package libs

import (
	"errors"
	"net/http"

	"github.com/ant-libs-go/util/logs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// TODO Can chagne to your `user` struct
type User struct {
	Id int32
}

type Context struct {
	ctx    *gin.Context
	role   Role
	Log    *logs.SessLog
	params Params
	user   *User
}

func NewContext(ctx *gin.Context, role Role, params Params) (r *Context, isAllow bool) {
	r = &Context{
		ctx:    ctx,
		role:   role,
		Log:    logs.New(uuid.NewV4().String()),
		params: params,
	}
	r.ctx.Set("logger", r.Log)

	defer func() {
		r.user = r.GetUser()
		// request log
		query := r.ctx.Request.URL.RawQuery
		if query != "" {
			query = "?" + query
		}
		r.Log.Infof("%s | %s %s %+v", r.ctx.ClientIP(), r.ctx.Request.Method, r.ctx.Request.URL.Path+query, params)
	}()

	// validate params
	if err := r.ctx.ShouldBind(params); err != nil {
		r.Log.Debugf("Params validate error: %s", err)
		r.RenderError(HttpStatusParamsErr, errors.New("Params validate error"))
		return
	}

	// validate auth
	if r.role == RoleGuest && r.GetUser() != nil {
		r.Log.Debugf("Auth validate error: already login")
		r.RenderError(HttpStatusAuthErr, errors.New("Auth Validate error"))
		return
	}
	if r.role == RoleAuth && r.GetUser() == nil {
		r.Log.Debugf("Auth validate error: not login")
		r.RenderError(HttpStatusAuthErr, errors.New("Auth Validate error"))
		return
	}
	isAllow = true
	return
}

func (this *Context) GetUser() *User {
	if this.user == nil {
		// TODO use token get user info
		/*
			r, err := GetUser(this.params.GetToken())
			if err != nil {
				this.Log.Warnf("get user error: %s", err.Error())
				return nil
			}
			this.user = r
		*/
	}
	return this.user
}

func (this *Context) GetUserId() (r int32) {
	if this.GetUser() != nil {
		r = this.GetUser().Id
	}
	return
}

func (this *Context) GetParams() Params {
	return this.params
}

/************ Render **************/
func (this *Context) Render(data interface{}) {
	h := gin.H{
		"code": HttpStatusOk,
		"data": data}
	this.Log.Infof("Response: %+v", h)
	this.ctx.JSON(http.StatusOK, h)
}

func (this *Context) RenderError(code HttpStatus, e error) {
	h := gin.H{
		"code": code,
		"msg":  e.Error()}
	this.Log.Infof("Response: %+v", h)
	this.ctx.JSON(http.StatusOK, h)
}

func (this *Context) RenderRealError(code HttpStatus, userE error, realE error) {
	if realE != nil {
		this.Log.Warnf(realE.Error())
	}
	if userE == nil {
		userE = errors.New("Server exception")
	}
	h := gin.H{
		"code": code,
		"msg":  userE.Error()}
	this.Log.Infof("Response: %+v", h)
	this.ctx.JSON(http.StatusOK, h)
}
