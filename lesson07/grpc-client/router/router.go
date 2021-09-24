package router

import (
	"github.com/gin-gonic/gin"
	"lesson07/grpc-client/controller"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/v1")
	todolist := apiV1.Group("/todolist")
	{
		todolist.POST("/add", controller.CreateToDoList)
		todolist.GET("/all", controller.ReadToDoList)
		todolist.POST("/del", controller.DeleteToDoList)
	}
	return r
}
