package client

import (
	"google.golang.org/grpc"
	"log"
)

func NewToDoClient () *grpc.ClientConn{
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
