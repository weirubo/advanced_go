package main

import (
	proto "lesson1/proto"
	"net"
	"net/http"
	"net/rpc"
)

// 服务器端（调用提供方）

func main() {
	_ = rpc.Register(new(proto.User))
	rpc.HandleHTTP()
	listener, _ := net.Listen("tcp", ":8081")
	_ = http.Serve(listener, nil)
}
