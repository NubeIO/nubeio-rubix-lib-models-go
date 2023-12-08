package model

import "gorm.io/datatypes"

type EmailConfigType string

const (
	EmailConfigTypeSmtp EmailConfigType = "SMTP"
	EmailConfigTypeSes  EmailConfigType = "SES"
)

type EmailConfig struct {
	CommonUUID
	Type      EmailConfigType `json:"type" gorm:"default:SMTP"`
	Enable    *bool           `json:"enable" gorm:"default:true"`
	QueueSize int             `json:"queue_size" gorm:"default:10"`
	Config    datatypes.JSON  `json:"config"`
}
