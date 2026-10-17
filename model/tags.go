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

// The many2many join tables below are declared so AutoMigrate can index them by entity; their primary keys lead with
// tag_tag, which leaves lookups by network, device or point scanning the whole table.
// The many2many relations create the tables, so migrate these after Network, Device, Point and Tag.

type NetworksTag struct {
	TagTag      string `gorm:"type:varchar(255);primaryKey"`
	NetworkUUID string `gorm:"type:varchar(255);primaryKey;index:idx_networks_tags_network_uuid"`
}

type DevicesTag struct {
	TagTag     string `gorm:"type:varchar(255);primaryKey"`
	DeviceUUID string `gorm:"type:varchar(255);primaryKey;index:idx_devices_tags_device_uuid"`
}

type PointsTag struct {
	TagTag    string `gorm:"type:varchar(255);primaryKey"`
	PointUUID string `gorm:"type:varchar(255);primaryKey;index:idx_points_tags_point_uuid"`
}
