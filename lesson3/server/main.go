package main

import (
	"context"
	"io"
	pb "lesson3/proto"
	"net"

	"google.golang.org/grpc"
)

// service
type UserService struct{}

func (u *UserService) GetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	userMap := map[int64]pb.UserResponse{
		1: {Uid: 1, Username: "frank", Age: 18, Wechat: "frank88"},
		2: {Uid: 2, Username: "lucy", Age: 17, Wechat: "lucy88"},
	}
	if data, ok := userMap[request.Uid]; ok {
		return &data, nil
	}
	response := &pb.UserResponse{
		Uid:      0,
		Username: "Unknow",
		Age:      0,
		Wechat:   "",
	}
	return response, nil
}

func (u *UserService) GetUserStream(stream pb.UserService_GetUserStreamServer) error {
	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		uid := request.GetUid()
		username := request.GetUsername()
		response := &pb.UserResponse{
			Uid:      uid,
			Username: username,
		}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, new(UserService))
	listener, _ := net.Listen("tcp", ":8081")
	_ = server.Serve(listener)
}
