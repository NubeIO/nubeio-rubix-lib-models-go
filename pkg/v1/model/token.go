package model

type Token struct {
	Name    string `json:"name" binding:"required" gorm:"type:varchar(255);unique;"`
	Token   string `json:"token" binding:"required" gorm:"type:varchar;unique;not null;default:null;primaryKey"`
	Blocked bool   `json:"blocked" gorm:"type:bool;not null;default:false;"`
}
