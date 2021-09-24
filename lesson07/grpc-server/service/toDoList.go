package service

import (
	"context"
	"lesson07/grpc-server/dao"
	pb "lesson07/grpc-server/pb/todoPb"
	"log"
)

type ToDo struct {
	pb.UnimplementedToDoListServer
}

func (t *ToDo) CreateToDo(ctx context.Context, in *pb.ToDoContent) (*pb.Result, error) {
	log.Printf("id: %d content:%v datetime:%d\n", in.GetId(), in.GetContent(), in.GetDatetime())
	record, err := dao.Add(ctx, in)
	data := &pb.Result{Record: record}
	return data, err
}
