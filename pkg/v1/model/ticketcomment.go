package model

type TicketComment struct {
	CommonUUID
	TicketUUID string `json:"ticket_uuid,omitempty" gorm:"type:varchar(255) references tickets;not null;"`
	Owner      string `json:"owner"`
	Content    string `json:"content"`
	CommonCreated
}
