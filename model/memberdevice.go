package model

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

type MemberDevice struct {
	CommonUUID
	MemberUUID string                        `json:"member_uuid,omitempty" gorm:"type:varchar(255) references members;not null;default:null;uniqueIndex:idx_member_devices_member_uuid_device_id"`
	DeviceID   string                        `json:"device_id" gorm:"uniqueIndex:idx_member_devices_member_uuid_device_id"`
	DeviceName *string                       `json:"device_name"`
	Platform   datatype.MemberDevicePlatform `json:"platform"`
}
