package main

import (
	"github.com/gin-gonic/gin"
	"qi.xiao/ginessential/controller"
	"qi.xiao/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //添加中间件保护用户信息接口
	return r
}
