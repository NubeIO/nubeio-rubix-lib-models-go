package model

import "time"

type AlertTransaction struct {
	CommonUUID
	AlertUUID string     `json:"alert_uuid" gorm:"type:varchar(255) references alerts;not null;default:null;primaryKey"`
	Status    string     `json:"status"`   // Active
	Severity  string     `json:"severity"` // Crucial
	Target    string     `json:"target,omitempty"`
	Title     string     `json:"title,omitempty"`
	Body      string     `json:"body,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
