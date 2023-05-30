package model

type ViewTemplateWidgetPointer struct {
	CommonUUID
	ViewTemplateWidgetUUID string `json:"view_template_widget_uuid,omitempty"  gorm:"type:varchar(255) references view_template_widgets;"`
	ViewUUID               string `json:"view_uuid,omitempty" gorm:"uniqueIndex:idx_view_template_widget_pointers_view_uuid_host_uuid_point_uuid;"`
	HostUUID               string `json:"host_uuid,omitempty"  gorm:"uniqueIndex:idx_view_template_widget_pointers_view_uuid_host_uuid_point_uuid;"`
	DeviceUUID             string `json:"device_uuid,omitempty"`
	PointUUID              string `json:"point_uuid,omitempty"  gorm:"uniqueIndex:idx_view_template_widget_pointers_view_uuid_host_uuid_point_uuid;"`
}
