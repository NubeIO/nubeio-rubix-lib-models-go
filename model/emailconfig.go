package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"gorm.io/datatypes"
)

type EmailConfig struct {
	CommonUUID
	Type      datatype.EmailConfigType `json:"type" gorm:"default:SMTP"`
	Enable    *bool                    `json:"enable" gorm:"default:true"`
	QueueSize int                      `json:"queue_size" gorm:"default:10"`
	Config    datatypes.JSON           `json:"config"`
}
