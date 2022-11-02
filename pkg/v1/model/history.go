package model

import "time"

// History for storing the all history
type History struct {
	HistoryID int       `json:"history_id" gorm:"primaryKey"`
	ID        int       `json:"id" gorm:"uniqueIndex:idx_histories_id_uuid_value_timestamp"`
	UUID      string    `json:"uuid" gorm:"uniqueIndex:idx_histories_id_uuid_value_timestamp"`
	Value     float64   `json:"value" gorm:"uniqueIndex:idx_histories_id_uuid_value_timestamp"`
	Timestamp time.Time `json:"timestamp" gorm:"uniqueIndex:idx_histories_id_uuid_value_timestamp"`
}
