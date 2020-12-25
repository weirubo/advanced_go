package main

import (
	"context"
	"fmt"
	"io"
	pb "lesson02/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

// client

func main() {
	conn, _ := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	// request := &pb.UserRequest{
	// 	Uid:      1,
	// 	Username: "frank",
	// }
	// response, _ := client.GetUser(context.Background(), request)
	// fmt.Printf("response:%+v\n", response)
	stream, err := client.GetUserStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			request := &pb.UserRequest{
				Uid:      1,
				Username: "lucy",
			}
			if err := stream.Send(request); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(response.GetUid())
		fmt.Println(response.GetUsername())
	}
}
