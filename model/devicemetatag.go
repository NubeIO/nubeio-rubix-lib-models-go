package model

type DeviceMetaTag struct {
	DeviceUUID string `json:"device_uuid,omitempty" gorm:"type:varchar(255) references devices;not null;default:null;primaryKey"`
	Key        string `json:"key,omitempty" gorm:"primaryKey"`
	Value      string `json:"value,omitempty"`
}
