package model

import "time"

type HostComment struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	HostUUID  string    `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;not null;"`
	Comment   string    `json:"comment"`
	UpdatedAt time.Time `json:"date,omitempty"`
}
