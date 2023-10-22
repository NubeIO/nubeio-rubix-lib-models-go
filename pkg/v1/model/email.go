package model

import "time"

type EmailStatus string

const (
	EmailStatusSending EmailStatus = "Sending"
	EmailStatusSent    EmailStatus = "Sent"
	EmailStatusFailed  EmailStatus = "Failed"
)

type Email struct {
	CommonUUID
	To          string             `json:"to"` // ; seperated emails
	Subject     string             `json:"subject"`
	Body        string             `json:"body"`
	Status      EmailStatus        `json:"status"`
	Error       *string            `json:"error"`
	CreatedAt   *time.Time         `json:"created_on,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_on,omitempty"`
	Attachments []*EmailAttachment `json:"attachments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
