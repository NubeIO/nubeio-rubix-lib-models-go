package model

import "github.com/shopspring/decimal"

type Group struct {
	UUID         string           `json:"uuid" gorm:"primary_key"`
	Name         string           `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_networks_name_location_uuid"`
	LocationUUID string           `json:"location_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES locations;not null;default:null;uniqueIndex:idx_networks_name_location_uuid"`
	Description  string           `json:"description"`
	Address      *string          `json:"address"`
	City         *string          `json:"city"`
	State        *string          `json:"state"`
	Zip          *int             `json:"zip"`
	Country      *string          `json:"country"`
	Lat          *decimal.Decimal `json:"lat"`
	Lon          *decimal.Decimal `json:"lon"`
	TimeZone     *string          `json:"time_zone"`
	Hosts        []*Host          `json:"hosts" gorm:"constraint:OnDelete:CASCADE"`
	Members      []*Member        `json:"members,omitempty" gorm:"many2many:member_groups;constraint:OnDelete:CASCADE"`
}
