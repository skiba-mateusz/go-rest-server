package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skiba-mateusz/go-rest-server/config"
	"github.com/skiba-mateusz/go-rest-server/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	cfg         config.Config
	mongoClient *mongo.Client
}

func NewAPIServer(cfg config.Config, mongoClient *mongo.Client) *APIServer {
	return &APIServer{
		cfg:         cfg,
		mongoClient: mongoClient,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	handler := handlers.New(s.cfg, s.mongoClient)

	mux.HandleFunc("POST /records", handler.GetRecords)
	mux.HandleFunc("POST /memory", handler.InsertRecord)
	mux.HandleFunc("GET /memory", handler.FindRecord)

	log.Println("Server is listening on port:", s.cfg.Port)

	return http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.Port), mux)
}
