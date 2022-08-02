package model

import "time"

// HistoryPostgresLog for storing the all history postgres logs
type HistoryPostgresLog struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UUID      string    `json:"uuid" gorm:"primary_key"`
	Value     float64   `json:"value" gorm:"primary_key"`
	Timestamp time.Time `json:"timestamp" gorm:"primary_key"`
}
