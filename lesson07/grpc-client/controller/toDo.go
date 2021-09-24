package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"lesson07/grpc-client/client"
	pb "lesson07/grpc-server/pb"
	"log"
	"net/http"
	"time"
)

func CreateToDo (ctx *gin.Context) {
	param := &pb.ToDoContent{
		Id:                   1,
		Content:              "看书一小时",
		Datetime:             time.Now().Unix(),
	}
	conn := client.NewToDoClient()
	defer conn.Close()
	cli := pb.NewToDoClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := cli.CreateToDo(ctx1, param)
	if err != nil {
		log.Fatalf("create todo err: %v", err)
	}
	log.Printf("todo content: %v", res.GetContent())
	data := pb.ToDoDetail{
		Id:                   res.Id,
		Content:              res.Content,
		Datetime:             res.Datetime,
		Created:              res.Created,
		Updated:              res.Updated,
	}
	ctx.JSON(http.StatusOK, data)
}
