package model

type TicketTeam struct {
	TicketUUID string `json:"ticket_uuid" gorm:"type:varchar(255) references tickets;not null;default:null;primaryKey"`
	TeamUUID   string `json:"team_uuid" gorm:"type:varchar(255) references teams;not null;default:null;primaryKey"`
}
