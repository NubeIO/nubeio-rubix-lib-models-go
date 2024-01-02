package model

type AlertMetaTag struct {
	AlertUUID string `json:"alert_uuid,omitempty" gorm:"type:varchar(255) references alerts;not null;default:null;primaryKey"`
	Key       string `json:"key,omitempty" gorm:"primaryKey"`
	Value     string `json:"value,omitempty"`
}
