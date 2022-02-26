package main

import (
	"github.com/weirubo/advanced_go/lesson08/grpc-server/dao"
	pb "github.com/weirubo/advanced_go/lesson08/grpc-server/pb/todoPb"
	"github.com/weirubo/advanced_go/lesson08/grpc-server/service"
	"google.golang.org/grpc"
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
