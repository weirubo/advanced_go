package main

import (
	"fmt"
	proto "lesson1/proto"
	"net/rpc"
)

// 客户端（调用方）

func main() {
	client, _ := rpc.DialHTTP("tcp", ":8081")
	id := 2
	var user proto.User
	// _ = client.Call("User.GetUser", id, &user)
	// fmt.Println(user)
	userCall := client.Go("User.GetUser", id, &user, nil)
	if replyCall := <-userCall.Done; replyCall != nil {
		fmt.Println(user)
	}
}
