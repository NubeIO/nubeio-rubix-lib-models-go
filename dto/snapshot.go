package dto

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type Snapshots struct {
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}
type CreateSnapshot struct {
	Description string `json:"description"`
}

type RestoreSnapshot struct {
	File        string `json:"file"`
	Description string `json:"description"`
}

type SnapshotStatus struct {
	CreateStatus  datatype.CreateStatus  `json:"create_status"`
	RestoreStatus datatype.RestoreStatus `json:"restore_status"`
}
