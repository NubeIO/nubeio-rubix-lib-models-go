package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type AlertTransaction struct {
	CommonUUID
	AlertUUID string                 `json:"alert_uuid" gorm:"type:varchar(255) references alerts;not null;default:null;primaryKey"`
	Status    datatype.AlertStatus   `json:"status"`   // Active
	Severity  datatype.AlertSeverity `json:"severity"` // Crucial
	Target    datatype.AlertTarget   `json:"target,omitempty"`
	Title     string                 `json:"title,omitempty"`
	Body      string                 `json:"body,omitempty"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
}
