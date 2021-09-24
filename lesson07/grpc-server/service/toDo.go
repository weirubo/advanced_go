package service

import (
	"context"
	pb "lesson07/grpc-server/pb"
	"log"
)

type ToDo struct {

}

func (t *ToDo) CreateToDo(ctx context.Context, in *pb.ToDoContent) (*pb.ToDoDetail, error) {
	log.Printf("id: %d content:%v datetime:%d\n", in.GetId(), in.GetContent(), in.GetDatetime())
	return &pb.ToDoDetail{
		Id:                   in.GetId(),
		Content:              in.GetContent(),
		Datetime:             in.GetDatetime(),
	}, nil
}

