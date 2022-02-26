package dao

import (
	"context"
	pb "github.com/weirubo/advanced_go/lesson08/grpc-server/pb/todoPb"
)

func Add(ctx context.Context, in *pb.ToDoListDetail) (int64, error) {
	err := eg.Sync2(new(pb.ToDoListDetail))
	if err != nil {
		return 0, err
	}
	toDoListDetail := new(pb.ToDoListDetail)
	toDoListDetail.Content = in.GetContent()
	toDoListDetail.Datetime = in.GetDatetime()
	rst, err := eg.Insert(toDoListDetail)
	return rst, err
}

func Select(ctx context.Context, in *pb.ToDoListPage) ([]*pb.ToDoListDetail, error) {
	toDoListDetails := make([]*pb.ToDoListDetail, 0)
	offset := (in.GetPage() - 1) * in.GetCount()
	err := eg.Limit(int(in.GetCount()), int(offset)).Find(&toDoListDetails)
	if err != nil {
		return nil, err
	}
	return toDoListDetails, nil
}
