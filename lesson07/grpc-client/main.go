package main

import (
	"grpc-demo/grpc-client/router"
)

func main () {
	engine := router.NewEngine()
	engine.Run(":8080")
}
