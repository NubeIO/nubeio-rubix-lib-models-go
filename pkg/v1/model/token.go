package model

type Token struct {
	Name    string `json:"name" binding:"required" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Token   string `json:"token" binding:"required" gorm:"type:varchar;unique;not null;default:null"`
	Blocked *bool  `json:"blocked" gorm:"not null;default:false"`
}
