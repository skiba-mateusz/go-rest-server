package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	log.Println("Server is listening on:", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, mux)
}
