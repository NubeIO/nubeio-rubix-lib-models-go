package model

type TeamView struct {
	TeamUUID string `json:"team_uuid" gorm:"type:varchar(255) references teams;not null;default:null;primaryKey"`
	ViewUUID string `json:"view_uuid" gorm:"type:varchar(255) references views;not null;default:null;primaryKey"`
}
