package main

import (
	"Desktop/advanced_go/lesson2/message"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"google.golang.org/protobuf/proto"
)

// server

type User struct{}

func (u *User) GetUser(request []byte, response *message.UserResponse) error {
	userMap := map[int64]message.UserResponse{
		1: {Uid: 1, Username: "frank", Age: 18, Gender: 1, Contact: &message.Contact{Tel: "13800138000", Wechat: "frank88", Email: "frank@gmail.com"}},
		2: {Uid: 2, Username: "lucy", Age: 17, Gender: 2, Contact: &message.Contact{Tel: "13800138666", Wechat: "lucy88", Email: "lucy@gmail.com"}},
	}

	// 反序列化
	var requestParam message.UserRequest
	err := proto.Unmarshal(request, &requestParam)
	if err != nil {
		log.Fatal(err)
	}

	if user, ok := userMap[requestParam.Uid]; ok {
		*response = user
	}

	return nil
}

func main() {
	_ = rpc.Register(new(User))
	rpc.HandleHTTP()
	listener, _ := net.Listen("tcp", ":8081")
	_ = http.Serve(listener, nil)
}
