package model

import "gorm.io/datatypes"

type View struct {
	CommonUUID
	Name             string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_views_name_location_uuid;uniqueIndex:idx_views_name_group_uuid;uniqueIndex:idx_views_name_host_uuid"`
	Type             string         `json:"type"`
	WidgetConfig     datatypes.JSON `json:"widget_config"`
	Theme            datatypes.JSON `json:"theme"`
	Data             string         `json:"data"`
	ViewTemplateUUID *string        `json:"view_template_uuid"`
	LocationUUID     *string        `json:"location_uuid,omitempty" gorm:"type:varchar(255) references locations;uniqueIndex:idx_views_name_location_uuid"`
	GroupUUID        *string        `json:"group_uuid,omitempty" gorm:"type:varchar(255) references groups;uniqueIndex:idx_views_name_group_uuid"`
	HostUUID         *string        `json:"host_uuid,omitempty" gorm:"type:varchar(255) references hosts;uniqueIndex:idx_views_name_host_uuid"`
	TeamViews        []*TeamView    `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Widgets          []*ViewWidget  `json:"widgets,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
