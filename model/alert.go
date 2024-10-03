package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type Alert struct {
	UUID               string                   `json:"uuid" gorm:"primarykey"`
	HostUUID           string                   `json:"host_uuid" gorm:"uniqueIndex:idx_alerts_host_uuid_entity_type_entity_uuid_type_title"`
	EntityType         datatype.AlertEntityType `json:"entity_type" gorm:"uniqueIndex:idx_alerts_host_uuid_entity_type_entity_uuid_type_title"`           // Device
	EntityUUID         string                   `json:"entity_uuid,omitempty" gorm:"uniqueIndex:idx_alerts_host_uuid_entity_type_entity_uuid_type_title"` // dev_abc123
	Type               datatype.AlertType       `json:"type" gorm:"uniqueIndex:idx_alerts_host_uuid_entity_type_entity_uuid_type_title"`                  // Ping
	Status             datatype.AlertStatus     `json:"status" gorm:"index"`                                                                              // Active
	Severity           datatype.AlertSeverity   `json:"severity"`                                                                                         // Crucial
	Target             datatype.AlertTarget     `json:"target,omitempty"`
	Title              string                   `json:"title,omitempty" gorm:"uniqueIndex:idx_alerts_host_uuid_entity_type_entity_uuid_type_title"`
	Body               string                   `json:"body,omitempty"`
	Source             *string                  `json:"source,omitempty"`
	NotifiedAt         *time.Time               `json:"notified_at,omitempty"`
	EmailedAt          *time.Time               `json:"emailed_at,omitempty"`
	AcknowledgeTimeout *time.Time               `json:"acknowledge_timeout,omitempty"`
	CreatedAt          *time.Time               `json:"created_at,omitempty"`
	LastUpdated        *time.Time               `json:"last_updated,omitempty"`
	Tags               []*Tag                   `json:"tags,omitempty" gorm:"many2many:alerts_tags;constraint:OnDelete:CASCADE"`
	MetaTags           []*AlertMetaTag          `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tickets            []*Ticket                `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams              []*AlertTeam             `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Transactions       []*AlertTransaction      `json:"transactions,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
