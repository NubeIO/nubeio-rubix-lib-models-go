package model

type Member struct {
	CommonUUID
	FirstName     *string         `json:"first_name"`
	LastName      *string         `json:"last_name"`
	Username      string          `json:"username" gorm:"type:varchar(255);unique;not null;default:null;"`
	Password      string          `json:"password"`
	Email         string          `json:"email"`
	Permission    *string         `json:"permission"`
	State         *string         `json:"state"`
	MemberDevices []*MemberDevice `json:"member_devices,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams         []*Team         `json:"-" gorm:"many2many:team_members;constraint:OnDelete:CASCADE"`
	Tickets       []*TicketMember `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

func (m *Member) MaskPassword() {
	m.Password = "******"
}

type MemberState string

const (
	Verified   MemberState = "VERIFIED"
	UnVerified MemberState = "UNVERIFIED"
)

var MemberStateMap = map[MemberState]int8{
	Verified:   0,
	UnVerified: 0,
}

type MemberPermission string

const (
	Read  MemberPermission = "READ"
	Write MemberPermission = "WRITE"
)

var MemberPermissionMap = map[MemberPermission]int8{
	Read:  0,
	Write: 0,
}
