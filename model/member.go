package model

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

type Member struct {
	CommonUUID
	FirstName     *string                   `json:"first_name"`
	LastName      *string                   `json:"last_name"`
	Username      string                    `json:"username" gorm:"type:varchar(255);unique;not null;default:null;"`
	Password      string                    `json:"password"`
	Email         string                    `json:"email"`
	Permission    datatype.MemberPermission `json:"permission"`
	State         datatype.MemberState      `json:"state"`
	MemberDevices []*MemberDevice           `json:"member_devices,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams         []*Team                   `json:"-" gorm:"many2many:team_members;constraint:OnDelete:CASCADE"`
	Tickets       []*TicketMember           `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

func (m *Member) MaskPassword() {
	m.Password = "******"
}
