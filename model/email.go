package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type Email struct {
	CommonUUID
	To          string               `json:"to"`  // ; seperated emails
	Cc          *string              `json:"cc"`  // ; seperated emails
	Bcc         *string              `json:"bcc"` // ; seperated emails
	Subject     string               `json:"subject"`
	Body        string               `json:"body"`
	Status      datatype.EmailStatus `json:"status"`
	Error       *string              `json:"error"`
	CreatedAt   *time.Time           `json:"created_on,omitempty"`
	UpdatedAt   *time.Time           `json:"updated_on,omitempty"`
	Attachments []*EmailAttachment   `json:"attachments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
