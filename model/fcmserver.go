package model

type FcmServer struct {
	UUID string `json:"-" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Key  string `json:"key"`
}
