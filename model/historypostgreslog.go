package model

import "time"

// HistoryPostgresLog for storing the all history postgres logs
type HistoryPostgresLog struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	PointUUID string    `json:"point_uuid"`
	HostUUID  string    `json:"host_uuid"`
	Value     *float64  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
