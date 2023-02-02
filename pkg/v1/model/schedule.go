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
	TimeZone          string         `json:"timezone"`
	IsActive          *bool          `json:"is_active"`
	ActiveWeekly      *bool          `json:"active_weekly"`
	ActiveException   *bool          `json:"active_exception"`
	ActiveEvent       *bool          `json:"active_event"`
	EnablePayload     bool           `json:"enable_payload"`
	MinPayload        float64        `json:"min_payload"`
	MaxPayload        float64        `json:"max_payload"`
	Payload           float64        `json:"payload"`
	PeriodStartString string         `json:"period_start_string"` // human readable timestamp
	PeriodStopString  string         `json:"period_stop_string"`  // human readable timestamp
	NextStartString   string         `json:"next_start_string"`   // human readable timestamp
	NextStopString    string         `json:"next_stop_string"`    // human readable timestamp
	Schedule          datatypes.JSON `json:"schedule"`
	CommonCreated
}

type ScheduleData struct {
	Schedules Schedules      `json:"schedules,omitempty"`
	Config    datatypes.JSON `json:"config,omitempty"`
}

type Schedules struct {
	Events    map[string]Events    `json:"events"`
	Weekly    map[string]Weekly    `json:"weekly"`
	Exception map[string]Exception `json:"exception"`
}

type Events struct {
	Name  string `json:"name"`
	Dates []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"dates"`
	Value  float64 `json:"value"`
	Color  string  `json:"color"`
	Enable *bool   `json:"enable"`
}

type Weekly struct {
	Name   string   `json:"name"`
	Days   []string `json:"days"`
	Start  string   `json:"start"`
	End    string   `json:"end"`
	Value  float64  `json:"value"`
	Color  string   `json:"color"`
	Enable *bool    `json:"enable"`
}

type Exception struct {
	Name  string `json:"name"`
	Dates []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"dates"`
	Value  float64 `json:"value"`
	Color  string  `json:"color"`
	Enable *bool   `json:"enable"`
}
