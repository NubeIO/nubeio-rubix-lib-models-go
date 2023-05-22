package model

import "time"

type RestoreStatus string

const (
	Restoring     RestoreStatus = "Restoring"
	Restored      RestoreStatus = "Restored"
	RestoreFailed RestoreStatus = "Failed"
)

type SnapshotRestoreLog struct {
	UUID        string        `json:"uuid" gorm:"primary_key" get:"true" delete:"true"`
	HostUUID    string        `json:"host_uuid" get:"true" post:"true" patch:"true"`
	Msg         string        `json:"msg" get:"true" post:"true" patch:"true"`
	Status      RestoreStatus `json:"status" get:"true" post:"true" patch:"true"`
	Description string        `json:"description" get:"true" post:"true" patch:"true"`
	CreatedAt   time.Time     `json:"created_at" get:"true"`
}
