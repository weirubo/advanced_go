package dao

import (
	"context"
	pb "lesson07/grpc-server/pb/todoPb"
)

func Add(ctx context.Context, in *pb.ToDoContent) (int64, error) {
	err := eg.Sync2(new(pb.ToDoDetail))
	if err != nil {
		return 0, err
	}
	toDo := new(pb.ToDoDetail)
	toDo.Content = in.GetContent()
	toDo.Datetime = in.GetDatetime()
	rst, err := eg.Insert(toDo)
	return rst, err
}
