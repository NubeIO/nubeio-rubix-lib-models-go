package model

import "gorm.io/datatypes"

type View struct {
	CommonUUID
	Name         string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_views_name_location_uuid;uniqueIndex:idx_views_name_group_uuid;uniqueIndex:idx_views_name_host_uuid"`
	LocationUUID *string        `json:"location_uuid,omitempty" gorm:"type:varchar(255) references locations;uniqueIndex:idx_views_name_location_uuid"`
	GroupUUID    *string        `json:"group_uuid,omitempty" gorm:"type:varchar(255) references groups;uniqueIndex:idx_views_name_group_uuid"`
	HostUUID     *string        `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;uniqueIndex:idx_views_name_host_uuid"`
	Description  *string        `json:"description"`
	Layout       datatypes.JSON `json:"layout"`
	TeamViews    []*TeamView    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
