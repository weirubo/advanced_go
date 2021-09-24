package main

import (
	"lesson07/grpc-client/router"
	"log"
)

const (
	port    = ":8080"
	address = "127.0.0.1" + port
)

func main() {
	engine := router.NewEngine()
	err := engine.Run(address)
	if err != nil {
		log.Fatalf("engine run error=%v", err)
	}
}
