package model

import "github.com/shopspring/decimal"

type Team struct {
	UUID string `json:"uuid" gorm:"primary_key"`
	CommonNameUnique
	Address  *string          `json:"address"`
	City     *string          `json:"city"`
	State    *string          `json:"state"`
	Zip      *int             `json:"zip"`
	Country  *string          `json:"country"`
	Lat      *decimal.Decimal `json:"lat"`
	Lon      *decimal.Decimal `json:"lon"`
	TimeZone *string          `json:"time_zone"`
	Members  []*Member        `json:"-" gorm:"many2many:member_teams;constraint:OnDelete:CASCADE"`
}
