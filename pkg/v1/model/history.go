package model

import "time"

// History for storing the all history
type History struct {
	HistoryID int       `json:"history_id" gorm:"autoIncrement;primaryKey"`
	ID        int       `json:"id" gorm:"uniqueIndex:idx_histories_id_point_uuid_host_uuid_value_timestamp"`
	PointUUID string    `json:"point_uuid" gorm:"uniqueIndex:idx_histories_id_point_uuid_host_uuid_value_timestamp"`
	HostUUID  string    `json:"host_uuid" gorm:"uniqueIndex:idx_histories_id_point_uuid_host_uuid_value_timestamp"`
	Value     *float64  `json:"value" gorm:"uniqueIndex:idx_histories_id_point_uuid_host_uuid_value_timestamp"`
	Timestamp time.Time `json:"timestamp" gorm:"uniqueIndex:idx_histories_id_point_uuid_host_uuid_value_timestamp"`
}
