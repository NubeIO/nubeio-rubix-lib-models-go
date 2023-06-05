package model

import "time"

// HistoryLog for storing the history logs
type HistoryLog struct {
	ID         int       `json:"id" gorm:"autoIncrement;primaryKey;index"`
	HostUUID   string    `json:"host_uuid"`
	LastSyncID int       `json:"last_sync_id"`
	Timestamp  time.Time `json:"timestamp"`
}
