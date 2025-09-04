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
	uc := NewUserServiceClient()

	mx.Handle("GET /authorize", handlers.NewAuthorize(port, uc))
	mx.Handle("GET /oauth-callback", handlers.NewOAuthCallback(uc))
	mx.Handle("POST /api/refresh-token", handlers.NewRefreshToken(uc))

	server := http.Server{
		Addr:    ":" + port,
		Handler: mx,
	}

	err := server.ListenAndServe()
	if err != nil {
		// TODO: log error
		l.Fatalf("server.ListenAndServe: %v", err)
	}
}
