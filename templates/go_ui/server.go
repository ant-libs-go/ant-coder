/* ######################################################################
# Author: (__AUTHOR__)
# Created Time: __CREATE_DATETIME__
# File Name: server.go
# Description:
####################################################################### */

package main

import (
	"__PROJECT_NAME__/controllers"
	"__PROJECT_NAME__/libs/middlewares"

	"github.com/gin-gonic/gin"
)

func NewUiServer() (r *gin.Engine, err error) {
	r = gin.New()
	r.Use(middlewares.Logger(), gin.Recovery())
	r.MaxMultipartMemory = 2 << 20 // 2M

	site := r.Group("")
	{
		controller := controllers.NewSiteController()
		site.GET("/ping", controller.Ping)
		site.HEAD("/ping", controller.Ping)
		site.POST("/login", controller.Login)
	}

	return
}
