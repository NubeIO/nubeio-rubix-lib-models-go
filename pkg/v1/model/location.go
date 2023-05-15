package model

type Location struct {
	CommonUUID
	CommonNameUnique
	Description string   `json:"description"`
	Groups      []*Group `json:"groups" gorm:"constraint:OnDelete:CASCADE"`
}
