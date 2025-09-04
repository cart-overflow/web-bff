package rpc

import (
	"log"

	"github.com/cart-overflow/user-api/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient(port string) pb.UserServiceClient {
	// TODO: secure
	conn, err := grpc.NewClient(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// TODO: logg error
		log.Fatalf("grpc.NewClient: %v", err)
	}

	return pb.NewUserServiceClient(conn)
}
