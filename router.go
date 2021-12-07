package main

import (
	"github.com/gin-gonic/gin"
	"qdgo/goweb/controller"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
