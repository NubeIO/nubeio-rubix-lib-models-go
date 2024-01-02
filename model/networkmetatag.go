package model

type NetworkMetaTag struct {
	NetworkUUID string `json:"network_uuid,omitempty" gorm:"type:varchar(255) references networks;not null;default:null;primaryKey"`
	Key         string `json:"key,omitempty" gorm:"primaryKey"`
	Value       string `json:"value,omitempty"`
}
