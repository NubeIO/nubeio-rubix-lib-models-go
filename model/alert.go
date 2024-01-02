package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type Alert struct {
	UUID         string                   `json:"uuid" gorm:"primarykey"`
	HostUUID     string                   `json:"host_uuid"`
	EntityType   datatype.AlertEntityType `json:"entity_type"`           // Device
	EntityUUID   string                   `json:"entity_uuid,omitempty"` // dev_abc123
	Type         datatype.AlertType       `json:"type"`                  // Ping
	Status       datatype.AlertStatus     `json:"status" gorm:"index"`   // Active
	Severity     datatype.AlertSeverity   `json:"severity"`              // Crucial
	Target       datatype.AlertTarget     `json:"target,omitempty"`
	Title        string                   `json:"title,omitempty"`
	Body         string                   `json:"body,omitempty"`
	Notified     *bool                    `json:"notified,omitempty" gorm:"default:false"`
	NotifiedAt   *time.Time               `json:"notified_at,omitempty"`
	CreatedAt    *time.Time               `json:"created_at,omitempty"`
	LastUpdated  *time.Time               `json:"last_updated,omitempty"`
	Tags         []*Tag                   `json:"tags,omitempty" gorm:"many2many:alerts_tags;constraint:OnDelete:CASCADE"`
	MetaTags     []*AlertMetaTag          `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tickets      []*Ticket                `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams        []*AlertTeam             `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Transactions []*AlertTransaction      `json:"transactions,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
