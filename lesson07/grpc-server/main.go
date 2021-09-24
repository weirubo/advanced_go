package main

import (
	"google.golang.org/grpc"
	pb "grpc-demo/grpc-server/pb"
	"grpc-demo/grpc-server/service"
	"log"
	"net"
)

const (
	port = ":8081"
)

func main () {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterToDoServer(server, new(service.ToDo))
	log.Printf("server listening at %v\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
