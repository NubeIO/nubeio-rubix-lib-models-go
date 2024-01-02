package model

type TeamMetaTag struct {
	TeamUUID string `json:"team_uuid,omitempty" gorm:"type:varchar(255) references teams;not null;default:null;primaryKey"`
	Key      string `json:"key,omitempty" gorm:"primaryKey"`
	Value    string `json:"value,omitempty"`
}
