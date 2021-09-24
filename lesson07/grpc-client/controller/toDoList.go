package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"lesson07/grpc-client/client"
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
	param := &pb.ToDoListContent{
		Id:       1,
		Content:  "看书一小时",
		Datetime: time.Now().Unix(),
	}
	conn := client.NewToDoListClient()
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
