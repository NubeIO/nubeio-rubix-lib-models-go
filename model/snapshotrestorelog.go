package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type SnapshotRestoreLog struct {
	UUID        string                 `json:"uuid" gorm:"primary_key" get:"true" delete:"true"`
	HostUUID    string                 `json:"host_uuid" get:"true" post:"true" patch:"true"`
	Msg         string                 `json:"msg" get:"true" post:"true" patch:"true"`
	Status      datatype.RestoreStatus `json:"status" get:"true" post:"true" patch:"true"`
	Description string                 `json:"description" get:"true" post:"true" patch:"true"`
	CreatedAt   time.Time              `json:"created_at" get:"true"`
}
