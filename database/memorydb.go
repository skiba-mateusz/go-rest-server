package database

import (
	"github.com/skiba-mateusz/go-rest-server/models"
)

type MemoryDB struct {
	records map[string]string
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		records: map[string]string{},
	}
}

func (d *MemoryDB) InsertRecord(record models.MemoryRecord) {
	d.records[record.Key] = record.Value
}

func (d *MemoryDB) FindRecord(key string) (models.MemoryRecord, bool) {
	if value, ok := d.records[key]; ok {
		return models.MemoryRecord{Key: key, Value: value}, true
	}
	return models.MemoryRecord{}, false
}
