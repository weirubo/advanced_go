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
		todolist.POST("/delete", controller.DeleteToDoList)
		todolist.POST("/edit", controller.UpdateToDoList)
		todolist.GET("/select", controller.ReadToDoList)
	}
	return r
}
