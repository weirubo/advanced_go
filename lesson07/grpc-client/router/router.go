package router

import (
	"github.com/gin-gonic/gin"
	"lesson07/grpc-client/controller"
)

func NewEngine ()  *gin.Engine{
	r := gin.Default()
	r.GET("/add", controller.CreateToDo)
	return r
}
