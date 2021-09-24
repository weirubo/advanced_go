package main

import (
	"lesson07/grpc-client/router"
)

func main () {
	engine := router.NewEngine()
	engine.Run(":8080")
}
