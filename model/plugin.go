package model

// Plugin holds information about the plugin.
type Plugin struct {
	UUID         string `json:"uuid" sql:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Name         string `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Enabled      bool   `json:"enabled"`
	MessageLevel string `json:"message_level"`
	Message      string `json:"message"`
	Config       []byte
	Storage      []byte
	Network      Network `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
}
