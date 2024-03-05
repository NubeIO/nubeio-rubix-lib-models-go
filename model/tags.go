package model

type Tag struct {
	Tag             string            `json:"tag" gorm:"type:varchar(255);unique;not null;default:null;primaryKey"`
	Networks        []*Network        `json:"networks,omitempty" gorm:"many2many:networks_tags;constraint:OnDelete:CASCADE"`
	Devices         []*Device         `json:"devices,omitempty" gorm:"many2many:devices_tags;constraint:OnDelete:CASCADE"`
	Points          []*Point          `json:"points,omitempty" gorm:"many2many:points_tags;constraint:OnDelete:CASCADE"`
	Alerts          []*Alert          `json:"alerts,omitempty" gorm:"many2many:alerts_tags;constraint:OnDelete:CASCADE"`
	Teams           []*Team           `json:"teams,omitempty" gorm:"many2many:teams_tags;constraint:OnDelete:CASCADE"`
	Hosts           []*Host           `json:"hosts,omitempty" gorm:"many2many:hosts_tags;constraint:OnDelete:CASCADE"`
	AlertConditions []*AlertCondition `json:"alert_conditions,omitempty" gorm:"many2many:alert_conditions_tags;constraint:OnDelete:CASCADE"`
}
