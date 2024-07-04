package models

import (
	"fmt"
	"time"
)

type RequestPayload struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

func (p *RequestPayload) Validate() error {
	if p.StartDate == "" || p.EndDate == "" {
		return fmt.Errorf("startDate and endDate cannot be empty")
	}
	if _, err := time.Parse("2006-01-02", p.StartDate); err != nil {
		return fmt.Errorf("invalid startDate")
	}
	if _, err := time.Parse("2006-01-02", p.EndDate); err != nil {
		return fmt.Errorf("invalid endDate")
	}
	if p.MinCount < 0 || p.MaxCount < 0 {
		return fmt.Errorf("minCount and maxCount must be positive numbers")
	}
	if p.MaxCount < p.MinCount {
		return fmt.Errorf("maxCount must be higher than minCount")
	}
	return nil
}

type MongoRecord struct {
	ID         string    `bson:"_id" json:"-"`
	Key        string    `bson:"key" json:"key"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	Count      []int     `bson:"count" json:"-"`
	Totalcount int       `bson:"totalCount" json:"totalCount"`
}

type ResponsePayload struct {
	Code    int           `json:"code"`
	Msg     string        `json:"msg"`
	Records []MongoRecord `json:"records"`
}

func ErrorResponsePayload(err error) ResponsePayload {
	return ResponsePayload{
		Code:    1,
		Msg:     err.Error(),
		Records: []MongoRecord{},
	}
}
