package main

import (
	"github.com/cart-overflow/web-bff/internal/rpc"
	"github.com/cart-overflow/web-bff/internal/server"
)

func main() {
	uc := rpc.NewUserServiceClient("6008")
	s := server.NewServer("4004", uc)
	s.Run()
}
