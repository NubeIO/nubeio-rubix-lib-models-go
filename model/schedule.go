package model

import (
	"gorm.io/datatypes"
)

type Schedule struct {
	CommonUUID
	CommonNameUnique
	CommonEnable
	CommonThingClass
	CommonThingType
	IsActive          *bool          `json:"is_active"`
	ActiveWeekly      *bool          `json:"active_weekly"`
	ActiveException   *bool          `json:"active_exception"`
	ActiveEvent       *bool          `json:"active_event"`
	EnablePayload     bool           `json:"enable_payload"`
	MinPayload        float64        `json:"min_payload"`
	MaxPayload        float64        `json:"max_payload"`
	Payload           float64        `json:"payload"`
	DefaultPayload    float64        `json:"default_payload"`
	PeriodStart       int64          `json:"period_start"`        // unix (seconds) timestamp
	PeriodStop        int64          `json:"period_stop"`         // unix (seconds) timestamp
	NextStart         int64          `json:"next_start"`          // unix (seconds) timestamp
	NextStop          int64          `json:"next_stop"`           // unix (seconds) timestamp
	PeriodStartString string         `json:"period_start_string"` // human readable timestamp
	PeriodStopString  string         `json:"period_stop_string"`  // human readable timestamp
	NextStartString   string         `json:"next_start_string"`   // human readable timestamp
	NextStopString    string         `json:"next_stop_string"`    // human readable timestamp
	Schedule          datatypes.JSON `json:"schedule"`
	CommonCreated
	GlobalUUID        string  `json:"global_uuid"`
	Connection        string  `json:"connection" gorm:"default:Connected"`
	ConnectionMessage *string `json:"connection_message" gorm:""`
	NullableOutput    *bool   `json:"nullable_output"`
}
