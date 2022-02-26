package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/weirubo/advanced_go/lesson08/grpc-server/pb/todoPb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const (
	grpcServerEndpoint = "localhost:8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterToDoListHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service endpoint: %v", err)
	}

	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}
}
