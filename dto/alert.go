package dto

import "time"

type AlertStatus struct {
	Status             string     `json:"status"`
	AcknowledgeTimeout *time.Time `json:"acknowledge_timeout"`
}
