package model

type TicketMember struct {
	TicketUUID string `json:"ticket_uuid" gorm:"type:varchar(255) references tickets;not null;default:null;primaryKey"`
	MemberUUID string `json:"member_uuid" gorm:"type:varchar(255) references members;not null;default:null;primaryKey"`
}
