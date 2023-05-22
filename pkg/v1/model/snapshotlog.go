package model

type SnapshotLog struct {
	File        string `json:"file" gorm:"primary_key" get:"true" delete:"true"`
	Description string `json:"description" get:"true" post:"true" patch:"true"`
}
