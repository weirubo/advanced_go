package main

import (
	"google.golang.org/grpc"
	"lesson07/grpc-server/dao"
	pb "lesson07/grpc-server/pb/todoPb"
	"lesson07/grpc-server/service"
	"log"
	"net"
)

const (
	port    = ":8081"
	address = "127.0.0.1" + port
)

func main() {
	InitEngine()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterToDoListServer(server, new(service.ToDoList))
	log.Printf("server listening at %v\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func InitEngine() {
	_, err := dao.NewEngine()
	if err != nil {
		log.Fatalf("mysql engine failed: %v", err)
	}
}
