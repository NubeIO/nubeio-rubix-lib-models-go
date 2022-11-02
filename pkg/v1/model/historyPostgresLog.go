package model

import "time"

// HistoryPostgresLog for storing the all history postgres logs
type HistoryPostgresLog struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	Value     float64   `json:"value" gorm:"primaryKey"`
	Timestamp time.Time `json:"timestamp" gorm:"primaryKey"`
}
