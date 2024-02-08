package model

import (
	"time"
)

type MetricLog struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;index"`
	Name      string    `json:"name,omitempty"`
	Value     *float64  `json:"value"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
