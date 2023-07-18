package model

import "gorm.io/datatypes"

type ViewWidget struct {
	CommonUUID
	Name         string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_view_widgets_name_view_uuid;"`
	Order        int            `json:"order"`
	X            int            `json:"x"`
	Y            int            `json:"y"`
	ViewUUID     string         `json:"view_uuid,omitempty" gorm:"type:varchar(255) references views;uniqueIndex:idx_view_widgets_name_view_uuid"`
	Type         *string        `json:"type"`
	Config       datatypes.JSON `json:"config"`
	HostUUID     string         `json:"host_uuid,omitempty"`
	Class        string         `json:"class,omitempty"`
	NetworkName  *string        `json:"network_name,omitempty"`
	DeviceUUID   *string        `json:"device_uuid,omitempty"`
	DeviceName   *string        `json:"device_name,omitempty"`
	PointUUID    *string        `json:"point_uuid,omitempty"`
	PointName    *string        `json:"point_name,omitempty"`
	ScheduleUUID *string        `json:"schedule_uuid,omitempty"`
	ScheduleName *string        `json:"schedule_name,omitempty"`
}

type Class string

const (
	ClassPoint    Class = "point"
	ClassSchedule Class = "schedule"
)

var ClassMap = map[Class]int8{
	ClassPoint:    0,
	ClassSchedule: 0,
}
