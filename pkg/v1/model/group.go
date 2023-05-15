package model

type Group struct {
	UUID         string  `json:"uuid" gorm:"primary_key"`
	Name         string  `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_networks_name_location_uuid"`
	LocationUUID string  `json:"location_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES locations;not null;default:null;uniqueIndex:idx_networks_name_location_uuid"`
	Description  string  `json:"description"`
	Hosts        []*Host `json:"hosts" gorm:"constraint:OnDelete:CASCADE"`
}
