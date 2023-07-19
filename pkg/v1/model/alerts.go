package model

import "time"

type AlertStatus string
type AlertEntityType string
type AlertType string
type AlertSeverity string
type AlertTarget string

const (
	AlertStatusActive       AlertStatus = "active"
	AlertStatusAcknowledged AlertStatus = "acknowledged"
	AlertStatusClosed       AlertStatus = "closed"
)

const (
	AlertEntityTypeGateway AlertEntityType = "gateway"
	AlertEntityTypeNetwork AlertEntityType = "network"
	AlertEntityTypeDevice  AlertEntityType = "device"
	AlertEntityTypePoint   AlertEntityType = "point"
	AlertEntityTypeService AlertEntityType = "service"
)

const (
	AlertTypePing      AlertType = "ping"
	AlertTypeFault     AlertType = "fault"
	AlertTypeThreshold AlertType = "threshold"
	AlertTypeFlatLine  AlertType = "flat-line"
)

const (
	AlertSeverityCrucial AlertSeverity = "crucial"
	AlertSeverityMinor   AlertSeverity = "minor"
	AlertSeverityInfo    AlertSeverity = "info"
	AlertSeverityWarning AlertSeverity = "warning"
)

const (
	AlertTargetMobile AlertTarget = "mobile"
	AlertTargetNone   AlertTarget = "none"
)

type Alert struct {
	UUID        string          `json:"uuid" gorm:"primarykey"`
	HostUUID    string          `json:"host_uuid"`
	EntityType  string          `json:"entity_type"`           // Device
	EntityUUID  string          `json:"entity_uuid,omitempty"` // dev_abc123
	Type        string          `json:"type"`                  // Ping
	Status      string          `json:"status"`                // Active
	Severity    string          `json:"severity"`              // Crucial
	Target      string          `json:"target,omitempty"`
	Title       string          `json:"title,omitempty"`
	Body        string          `json:"body,omitempty"`
	Notified    *bool           `json:"notified,omitempty" gorm:"default:false"`
	NotifiedAt  *time.Time      `json:"notified_at,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	LastUpdated *time.Time      `json:"last_updated,omitempty"`
	Tags        []*Tag          `json:"tags,omitempty" gorm:"many2many:alerts_tags;constraint:OnDelete:CASCADE"`
	MetaTags    []*AlertMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tickets     []*Ticket       `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams       []*AlertTeam    `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

type AlertMetaTag struct {
	AlertUUID string `json:"alert_uuid,omitempty" gorm:"type:varchar(255) references alerts;not null;default:null;primaryKey"`
	Key       string `json:"key,omitempty" gorm:"primaryKey"`
	Value     string `json:"value,omitempty"`
}

type AlertClosed struct {
	Alert
	ClosedAt *time.Time `json:"closed_at,omitempty"`
}
