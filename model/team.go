package model

import "github.com/shopspring/decimal"

type Team struct {
	UUID string `json:"uuid" gorm:"primary_key"`
	CommonNameUnique
	Address         *string               `json:"address"`
	City            *string               `json:"city"`
	State           *string               `json:"state"`
	Zip             *int                  `json:"zip"`
	Country         *string               `json:"country"`
	Lat             *decimal.Decimal      `json:"lat"`
	Lon             *decimal.Decimal      `json:"lon"`
	TimeZone        *string               `json:"time_zone"`
	Tags            []*Tag                `json:"tags,omitempty" gorm:"many2many:teams_tags;constraint:OnDelete:CASCADE"`
	MetaTags        []*TeamMetaTag        `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Members         []*Member             `json:"members,omitempty" gorm:"many2many:team_members;constraint:OnDelete:CASCADE"`
	Views           []*TeamView           `json:"views,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tickets         []*TicketTeam         `json:"tickets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Alerts          []*AlertTeam          `json:"alerts,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	AlertConditions []*AlertConditionTeam `json:"alert_conditions,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
