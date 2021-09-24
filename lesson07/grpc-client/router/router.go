package router

import (
	"github.com/gin-gonic/gin"
	"grpc-demo/grpc-client/controller"
)

func NewEngine ()  *gin.Engine{
	r := gin.Default()
	r.GET("/add", controller.CreateToDo)
	return r
}
