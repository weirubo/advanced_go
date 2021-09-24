package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "lesson07/grpc-client/pb/todoPb"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateToDoList(ctx *gin.Context) {
	param := new(pb.ToDoListContent)
	param.Datetime = time.Now().Unix()
	err := ctx.ShouldBind(param)
	if err != nil {
		log.Fatalf("param error=%+v", param)
	}
	conn := NewToDoListClient()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("conn close error=%v", err)
		}
	}()
	cli := pb.NewToDoListClient(conn)
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := cli.CreateToDoList(ctx1, param)
	if err != nil {
		log.Fatalf("create todolist err: %v", err)
	}
	rst := Result{
		Code:    10001,
		Message: "create todolist success",
		Data:    res.Record,
	}
	ctx.JSON(http.StatusOK, rst)
}

func DeleteToDoList(ctx *gin.Context) {

}

func UpdateToDoList(ctx *gin.Context) {

}

func ReadToDoList(ctx *gin.Context) {

}

func NewToDoListClient() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
