package model

type EmailAttachment struct {
	CommonUUID
	EmailUUID string `json:"email_uuid,omitempty" gorm:"type:varchar(255) references emails;not null;default:null;"`
	Path      string `json:"path"`
}
