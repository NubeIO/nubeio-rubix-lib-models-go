package model

import "time"

// History for storing the all history
type History struct {
	HistoryID int       `json:"history_id" gorm:"primaryKey"`
	ID        int       `json:"id" gorm:"uniqueIndex:id_uuid_value_timestamp_composite_index"`
	UUID      string    `json:"uuid" gorm:"uniqueIndex:id_uuid_value_timestamp_composite_index"`
	Value     float64   `json:"value" gorm:"uniqueIndex:id_uuid_value_timestamp_composite_index"`
	Timestamp time.Time `json:"timestamp" gorm:"uniqueIndex:id_uuid_value_timestamp_composite_index"`
}
