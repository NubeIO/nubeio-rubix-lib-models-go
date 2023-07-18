package model

type AlertTeam struct {
	AlertUUID string `json:"alert_uuid" gorm:"type:varchar(255) references alerts;not null;default:null;primaryKey"`
	TeamUUID  string `json:"team_uuid" gorm:"type:varchar(255) references teams;not null;default:null;primaryKey"`
}
