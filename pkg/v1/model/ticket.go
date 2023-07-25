package model

type TicketPriority string

const (
	TicketPriorityLow      TicketPriority = "LOW"
	TicketPriorityMedium   TicketPriority = "MEDIUM"
	TicketPriorityHigh     TicketPriority = "HIGH"
	TicketPriorityCritical TicketPriority = "CRITICAL"
)

type TicketStatus string

const (
	TicketStatusNew      TicketStatus = "NEW"
	TicketStatusReplied  TicketStatus = "REPLIED"
	TicketStatusResolved TicketStatus = "RESOLVED"
	TicketStatusClosed   TicketStatus = "CLOSED"
	TicketStatusBlocked  TicketStatus = "BLOCKED"
)

type Ticket struct {
	CommonUUID
	AlertUUID *string          `json:"alert_uuid,omitempty" gorm:"type:varchar(255) references alerts;"`
	Issuer    string           `json:"issuer"`
	Subject   string           `json:"subject"`
	Content   string           `json:"content"`
	Priority  TicketPriority   `json:"priority"`
	Status    TicketStatus     `json:"status"`
	Comments  []*TicketComment `json:"comments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams     []*TicketTeam    `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Members   []*TicketMember  `json:"members,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	CommonCreated
}
