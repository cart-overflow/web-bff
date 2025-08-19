package app

import (
	"log"
	"net/http"
	"os"
)

func RunServer() {
	l := log.New(os.Stdout, "web-bff: ", log.LstdFlags)
	mx := http.NewServeMux()

	mx.HandleFunc("/api/refresh-token", func(w http.ResponseWriter, r *http.Request) {
		l.Printf("Got request from %s", r.RemoteAddr)
		w.WriteHeader(200)
	})

	server := http.Server{
		Addr:    ":6000",
		Handler: mx,
	}

	err := server.ListenAndServe()
	if err != nil {
		l.Printf("TODO: handle ListenAndServe error")
	}
}
