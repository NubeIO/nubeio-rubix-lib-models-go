package model

type AlertConditionTeam struct {
	AlertConditionUUID string `json:"alert_condition_uuid" gorm:"type:varchar(255) references alert_conditions;not null;default:null;primaryKey"`
	TeamUUID           string `json:"team_uuid" gorm:"type:varchar(255) references teams;not null;default:null;primaryKey"`
}
