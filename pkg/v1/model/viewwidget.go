package model

import "gorm.io/datatypes"

type ViewWidget struct {
	CommonUUID
	Name        string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_view_widgets_name_view_uuid;"`
	Order       int            `json:"order"`
	X           int            `json:"x"`
	Y           int            `json:"y"`
	ViewUUID    string         `json:"view_uuid,omitempty" gorm:"type:varchar(255) references views;uniqueIndex:idx_view_widgets_name_view_uuid"`
	Type        *string        `json:"type"`
	Config      datatypes.JSON `json:"config"`
	NetworkName string         `json:"network_name"`
	DeviceName  string         `json:"device_name"`
	PointName   string         `json:"point_name"`
	HostUUID    string         `json:"host_uuid,omitempty" gorm:"uniqueIndex:idx_view_widgets_host_uuid_point_uuid"`
	DeviceUUID  string         `json:"device_uuid,omitempty"`
	PointUUID   string         `json:"point_uuid,omitempty" gorm:"uniqueIndex:idx_view_widgets_host_uuid_point_uuid"`
}
