package model

type User struct {
	Username string `json:"username" binding:"required" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Password string `json:"password" binding:"required"`
}
