package service

import (
	"context"
	"lesson07/grpc-server/dao"
	pb "lesson07/grpc-server/pb/todoPb"
	"log"
)

type ToDoList struct {
	pb.UnimplementedToDoListServer
}

func (t *ToDoList) CreateToDoList(ctx context.Context, in *pb.ToDoListDetail) (*pb.ToDoListResult, error) {
	log.Printf("id: %d content:%v datetime:%d\n", in.GetId(), in.GetContent(), in.GetDatetime())
	record, err := dao.Add(ctx, in)
	data := &pb.ToDoListResult{Record: record}
	return data, err
}

func (t *ToDoList) ReadToDoList(ctx context.Context, in *pb.ToDoListPage) (*pb.ToDoListByPage, error) {
	return nil, nil
}
