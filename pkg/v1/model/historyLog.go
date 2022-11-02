package model

import "time"

// HistoryLog for storing the history logs
type HistoryLog struct {
	ID                   int       `json:"id" gorm:"autoIncrement;primaryKey;index"`
	FlowNetworkCloneUUID string    `json:"flow_network_clone_uuid"`
	LastSyncID           int       `json:"last_sync_id"`
	Timestamp            time.Time `json:"timestamp"`
}
