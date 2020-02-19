/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: site_controller.go
# Description:
####################################################################### */

package controllers

import (
	"net/http"

	"__PROJECT_NAME__/libs"

	"github.com/gin-gonic/gin"
)

type SiteController struct {
	CommonController
}

func NewSiteController() *SiteController {
	o := &SiteController{}
	return o
}

func (this *SiteController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong, test ok")
}

func (this *SiteController) Login(c *gin.Context) {
	ctx, isAllow := libs.NewContext(c, libs.RoleAuth, &libs.LoginReq{})
	if !isAllow {
		return
	}

	p := ctx.GetParams().(*libs.LoginReq)
	r := libs.LoginResp{}

	// code in here

	ctx.Render(r)
}
