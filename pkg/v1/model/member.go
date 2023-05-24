package model

type Member struct {
	CommonUUID
	FirstName     *string         `json:"first_name"`
	LastName      *string         `json:"last_name"`
	Username      string          `json:"username" gorm:"type:varchar(255);unique;not null;default:null;"`
	Password      string          `json:"password"`
	Email         string          `json:"email"`
	State         *string         `json:"state"`
	MemberDevices []*MemberDevice `json:"member_devices,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Teams         []*Team         `json:"teams" gorm:"many2many:member_teams;constraint:OnDelete:CASCADE"`
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
