package server

import (
	"log"
	"net/http"

	"github.com/cart-overflow/user-api/pkg/pb"
)

type Server struct {
	port string
	serv *http.Server
	mx   *http.ServeMux
	uc   pb.UserServiceClient
}

func NewServer(port string, uc pb.UserServiceClient) *Server {
	mx := http.NewServeMux()
	serv := &http.Server{
		Addr:    ":" + port,
		Handler: mx,
	}

	return &Server{port, serv, mx, uc}
}

func (s *Server) Run() {
	s.mx.Handle("GET /authorize", NewAuthorize(s.port, s.uc))
	s.mx.Handle("GET /oauth-callback", NewOAuthCallback(s.uc))
	s.mx.Handle("POST /api/refresh-token", NewRefreshToken(s.uc))

	err := s.serv.ListenAndServe()
	if err != nil {
		// TODO: log error
		log.Fatalf("server.ListenAndServe: %v", err)
	}
}
