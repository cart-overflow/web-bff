package app

import (
	"log"

	"github.com/cart-overflow/user-api/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() pb.UserServiceClient {
	addr := ":6002"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient: %v", err)
	}

	return pb.NewUserServiceClient(conn)
}
