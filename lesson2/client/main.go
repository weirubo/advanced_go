package main

import (
	"Desktop/advanced_go/lesson2/message"
	"fmt"
	"log"
	"net/rpc"

	"github.com/golang/protobuf/proto"
)

// client

func main() {
	client, _ := rpc.DialHTTP("tcp", ":8081")
	request := &message.UserRequest{
		Uid:      1,
		Username: "frank",
	}
	// 序列化
	requestByte, err := proto.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var user message.UserResponse
	_ = client.Call("User.GetUser", requestByte, &user)
	fmt.Printf("user:%+v", user)
}
