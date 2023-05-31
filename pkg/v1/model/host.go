package model

import "time"

type Host struct {
	UUID           string         `json:"uuid" gorm:"primaryKey"`
	GlobalUUID     string         `json:"global_uuid"`
	DeviceType     string         `json:"device_type"`
	Name           string         `json:"name"  gorm:"type:varchar(255);not null;uniqueIndex:idx_hosts_name_group_uuid"`
	GroupUUID      string         `json:"group_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES groups;not null;default:null;uniqueIndex:idx_hosts_name_group_uuid"`
	Enable         *bool          `json:"enable"`
	Description    string         `json:"description"`
	IP             string         `json:"ip"`
	BiosPort       int            `json:"bios_port"`
	Port           int            `json:"port"`
	HTTPS          *bool          `json:"https"`
	IsOnline       *bool          `json:"is_online"`
	IsValidToken   *bool          `json:"is_valid_token"`
	ExternalToken  string         `json:"external_token"`
	VirtualIP      string         `json:"virtual_ip"`
	ReceivedBytes  int            `json:"received_bytes"`
	SentBytes      int            `json:"sent_bytes"`
	ConnectedSince string         `json:"connected_since"`
	Tags           []*HostTag     `json:"tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Comments       []*HostComment `json:"comments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Views          []*View        `json:"views" gorm:"constraint:OnDelete:CASCADE"`
}

type HostTag struct {
	HostUUID string `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;not null;default:null;primaryKey"`
	Tag      string `json:"tag" gorm:"primaryKey"`
}

type HostComment struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	HostUUID  string    `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;not null;"`
	Comment   string    `json:"comment"`
	UpdatedAt time.Time `json:"date,omitempty"`
}
