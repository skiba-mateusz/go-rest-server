package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/skiba-mateusz/go-rest-server/config"
	"github.com/skiba-mateusz/go-rest-server/database"
	"github.com/skiba-mateusz/go-rest-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	cfg         config.Config
	mongoClient *mongo.Client
	memoryDB    *database.MemoryDB
}

func New(cfg config.Config, mongoClient *mongo.Client) *Handler {
	return &Handler{
		cfg:         cfg,
		mongoClient: mongoClient,
		memoryDB:    database.NewMemoryDB(),
	}
}

func (h *Handler) GetRecords(w http.ResponseWriter, r *http.Request) {
	payload := models.RequestPayload{}
	err := parseJSON(r, &payload)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		writeJSON(w, http.StatusBadRequest, models.ErrorResponsePayload(err))
		return
	}

	err = payload.Validate()
	if err != nil {
		log.Printf("Error validating payload: %v", err)
		writeJSON(w, http.StatusBadRequest, models.ErrorResponsePayload(err))
		return
	}

	startDate, _ := time.Parse("2006-01-02", payload.StartDate)
	endDate, _ := time.Parse("2006-01-02", payload.EndDate)
	collection := h.mongoClient.Database(h.cfg.DBName).Collection("records")
	projectStage := bson.D{{
		Key: "$project", Value: bson.M{
			"key":       1,
			"createdAt": 1,
			"totalCount": bson.M{
				"$sum": "$count",
			},
		},
	}}
	matchStage := bson.D{{
		Key: "$match", Value: bson.M{
			"createdAt": bson.M{
				"$gte": startDate,
				"$lte": endDate,
			},
			"totalCount": bson.M{
				"$gte": payload.MinCount,
				"$lte": payload.MaxCount,
			},
		},
	}}
	sortStage := bson.D{{
		Key: "$sort", Value: bson.M{
			"createdAt": 1,
		},
	}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{projectStage, matchStage, sortStage})
	if err != nil {
		log.Printf("error aggregating collection, error: %v", err)
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponsePayload(err))
		return
	}
	defer cursor.Close(ctx)

	records := []models.MongoRecord{}
	if err = cursor.All(ctx, &records); err != nil {
		log.Printf("error reading data, error: %v", err)
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponsePayload(err))
		return
	}

	writeJSON(w, http.StatusOK, models.ResponsePayload{
		Code:    0,
		Msg:     "Success",
		Records: records,
	})
}

func (h *Handler) InsertRecord(w http.ResponseWriter, r *http.Request) {
	record := models.MemoryRecord{}
	if err := parseJSON(r, &record); err != nil {
		log.Printf("Invalid request body, error: %v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if record.Key == "" || record.Value == "" {
		err := fmt.Errorf("key and value are required")
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.memoryDB.InsertRecord(record)
	log.Printf("Record created: %v", record)
}

func (h *Handler) FindRecord(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		err := fmt.Errorf("key is empty")
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, ok := h.memoryDB.FindRecord(key)
	if !ok {
		err := fmt.Errorf("could not find record for provided key")
		log.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, record)
}
