package model

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

type Ticket struct {
	CommonUUID
	AlertUUID *string                 `json:"alert_uuid,omitempty" gorm:"type:varchar(255) references alerts;"`
	Issuer    string                  `json:"issuer"`
	Subject   string                  `json:"subject"`
	Content   string                  `json:"content"`
	Priority  datatype.TicketPriority `json:"priority"`
	Status    datatype.TicketStatus   `json:"status"`
	Comments  []*TicketComment        `json:"comments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams     []*TicketTeam           `json:"teams,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Members   []*TicketMember         `json:"members,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	CommonCreated
}
