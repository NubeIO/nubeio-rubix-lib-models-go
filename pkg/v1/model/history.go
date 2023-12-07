package model

import "time"

// History for storing the all history
type History struct {
	HistoryID int       `json:"history_id" gorm:"autoIncrement;primaryKey"`
	PointUUID string    `json:"point_uuid" gorm:"uniqueIndex:idx_histories_point_uuid_host_uuid_timestamp"`
	HostUUID  string    `json:"host_uuid" gorm:"uniqueIndex:idx_histories_point_uuid_host_uuid_timestamp"`
	Value     *float64  `json:"value"`
	Timestamp time.Time `json:"timestamp" gorm:"uniqueIndex:idx_histories_point_uuid_host_uuid_timestamp"`
}
