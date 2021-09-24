package dao

import (
	"context"
	pb "lesson07/grpc-server/pb/todoPb"
)

func Add(ctx context.Context, in *pb.ToDoListContent) (int64, error) {
	err := eg.Sync2(new(pb.ToDoListDetail))
	if err != nil {
		return 0, err
	}
	toDo := new(pb.ToDoListDetail)
	toDo.Content = in.GetContent()
	toDo.Datetime = in.GetDatetime()
	rst, err := eg.Insert(toDo)
	return rst, err
}
