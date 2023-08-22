package model

import "gorm.io/datatypes"

type ViewWidget struct {
	CommonUUID
	Name           string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_view_widgets_name_view_uuid;"`
	Order          int            `json:"order"`
	X              int            `json:"x"`
	Y              int            `json:"y"`
	ViewUUID       string         `json:"view_uuid,omitempty" gorm:"type:varchar(255) references views;uniqueIndex:idx_view_widgets_name_view_uuid"`
	Type           *string        `json:"type"`
	Config         datatypes.JSON `json:"config"`
	HostUUID       string         `json:"host_uuid,omitempty"`
	Class          string         `json:"class,omitempty"`
	HasDiffRW      *bool          `json:"has_diff_rw,omitempty"`
	NetworkUUID    *string        `json:"network_uuid,omitempty"` // rubix-ce needs this
	PointUUID      *string        `json:"point_uuid,omitempty"`
	WritePointUUID *string        `json:"write_point_uuid,omitempty"`
	ScheduleUUID   *string        `json:"schedule_uuid,omitempty"`
	Query          *string        `json:"query,omitempty"`
	// Computed values
	NetworkName      *string `json:"network_name,omitempty"`
	DeviceName       *string `json:"device_name,omitempty"`
	PointName        *string `json:"point_name,omitempty"`
	WriteNetworkName *string `json:"write_network_name,omitempty"`
	WriteDeviceName  *string `json:"write_device_name,omitempty"`
	WritePointName   *string `json:"write_point_name,omitempty"`
	ScheduleName     *string `json:"schedule_name,omitempty"`
}

type Class string

const (
	ClassPoint    Class = "point"
	ClassSchedule Class = "schedule"
	ClassQuery    Class = "query"
)

var ClassMap = map[Class]int8{
	ClassPoint:    0,
	ClassSchedule: 0,
	ClassQuery:    0,
}
