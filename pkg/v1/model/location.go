package model

type Location struct {
	CommonUUID
	CommonNameUnique
	Description string   `json:"description"`
	Groups      []*Group `json:"groups" gorm:"constraint:OnDelete:CASCADE"`
	Views       []*View  `json:"views" gorm:"constraint:OnDelete:CASCADE"`
}
