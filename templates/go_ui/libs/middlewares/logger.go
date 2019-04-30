/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-13 10:10:27
# File Name: logger.go
# Description:
####################################################################### */

package middlewares

import (
	"time"

	"github.com/ant-libs-go/util/logs"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		// Process request
		ctx.Next()
		// Stop timer
		end := time.Now()

		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		if query != "" {
			path = path + "?" + query
		}
		//fmt.Println("%s | %s %s", ctx.ClientIP(), ctx.Request.Method, path)

		if v, exists := ctx.Get("logger"); exists {
			v.(*logs.SessLog).Infof("code: %d, take: %s", ctx.Writer.Status(), end.Sub(start))
		}

		/*
			ctx.MustGet("logger").(*log.SessLog).Infof("%s | %s %s %d %s",
				ctx.ClientIP(),
				ctx.Request.Method,
				path,
				ctx.Writer.Status(),
				end.Sub(start))
		*/
	}
}
