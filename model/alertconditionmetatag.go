package model

type AlertConditionMetaTag struct {
	AlertConditionUUID string `json:"alert_condition_uuid,omitempty" gorm:"type:varchar(255) references alert_conditions;not null;default:null;primaryKey"`
	Key                string `json:"key,omitempty" gorm:"primaryKey"`
	Value              string `json:"value,omitempty"`
}
