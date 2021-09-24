package router

import (
	"github.com/gin-gonic/gin"
	"lesson07/grpc-client/controller"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/add", controller.CreateToDoList)
	}
	return r
}
