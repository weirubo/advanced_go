package main

import (
	pb "Desktop/advanced_go/lesson2/proto"
	"fmt"
	"log"
	"net/rpc"

	"github.com/golang/protobuf/proto"
)

// client

func main() {
	client, _ := rpc.DialHTTP("tcp", ":8081")
	request := &pb.UserRequest{
		Uid:      1,
		Username: "frank",
	}
	// 序列化
	requestByte, err := proto.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var user pb.UserResponse
	_ = client.Call("User.GetUser", requestByte, &user)
	fmt.Printf("user:%+v", user)
}
