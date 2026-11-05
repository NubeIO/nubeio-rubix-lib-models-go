package model

type HostMetaTag struct {
	HostUUID string `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;not null;default:null;primaryKey"`
	Key      string `json:"key,omitempty" gorm:"primaryKey"`
	Value    string `json:"value,omitempty"`
}
