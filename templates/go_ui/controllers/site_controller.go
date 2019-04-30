/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-15 08:34:12
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

func (this *SiteController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong, test ok")
}

func (this *SiteController) Login(c *gin.Context) {
	p := &libs.LoginParams{}
	ctx, isAllow := libs.NewContext(c, libs.RoleAuth, p)
	if !isAllow {
		return
	}

	r := libs.LoginParams{}

	// code in here

	ctx.Render(r)
}
