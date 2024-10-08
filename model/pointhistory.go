package model

import (
	"time"
)

type PointHistory struct {
	ID        int       `json:"id" gorm:"autoIncrement;primaryKey;index"`
	PointUUID string    `json:"point_uuid,omitempty" gorm:"type:varchar(255);not null;default:null;uniqueIndex:idx_point_histories_point_uuid_timestamp"`
	Value     *float64  `json:"present_value"`
	Timestamp time.Time `json:"timestamp,omitempty" gorm:"uniqueIndex:idx_point_histories_point_uuid_timestamp"`
}
