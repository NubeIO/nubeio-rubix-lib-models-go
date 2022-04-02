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
	IsActive *bool          `json:"is_active"`
	IsGlobal *bool          `json:"is_global"`
	Schedule datatypes.JSON `json:"schedule"`
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
	Value int    `json:"value"`
	Color string `json:"color"`
}

type Weekly struct {
	Name  string   `json:"name"`
	Days  []string `json:"days"`
	Start string   `json:"start"`
	End   string   `json:"end"`
	Value int      `json:"value"`
	Color string   `json:"color"`
}

type Exception struct {
	Name  string `json:"name"`
	Dates []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"dates"`
	Value int    `json:"value"`
	Color string `json:"color"`
}
