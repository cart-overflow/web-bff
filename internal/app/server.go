package app

import (
	"log"
	"net/http"
	"os"

	"github.com/cart-overflow/web-bff/internal/handlers"
)

func RunServer() {
	port := "6001"
	l := log.New(os.Stdout, "web-bff: ", log.LstdFlags)
	mx := http.NewServeMux()
	userClient := NewUserServiceClient()

	mx.Handle("GET /authorize", handlers.NewAuthorize(port, userClient))
	mx.Handle("GET /oauth-callback", handlers.NewOAuthCallback())
	mx.Handle("POST /api/refresh-token", handlers.NewRefreshToken())

	server := http.Server{
		Addr:    ":" + port,
		Handler: mx,
	}

	err := server.ListenAndServe()
	if err != nil {
		l.Fatalf("server.ListenAndServe: %v", err)
	}
}
