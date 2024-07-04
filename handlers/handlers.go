package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/skiba-mateusz/go-rest-server/config"
	"github.com/skiba-mateusz/go-rest-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	cfg         config.Config
	mongoClient *mongo.Client
}

func New(cfg config.Config, mongoClient *mongo.Client) *Handler {
	return &Handler{
		cfg:         cfg,
		mongoClient: mongoClient,
	}
}

func (h *Handler) GetRecords(w http.ResponseWriter, r *http.Request) {
	payload := models.RequestPayload{}
	err := parseJSON(r, &payload)
	if err != nil {
		log.Printf("Error parsing JSON: %v", error)
		writeJSON(w, http.StatusBadRequest, models.ErrorResponsePayload(err))
		return
	}

	err = payload.Validate()
	if err != nil {
		log.Printf("Error validating payload: %v", error)
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
