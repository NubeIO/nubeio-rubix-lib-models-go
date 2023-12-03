package model

type ViewTemplateWidgetPointer struct {
	CommonUUID
	ViewTemplateWidgetUUID string  `json:"view_template_widget_uuid,omitempty"  gorm:"type:varchar(255) references view_template_widgets;"`
	ViewUUID               string  `json:"view_uuid,omitempty"`
	HostUUID               string  `json:"host_uuid,omitempty"`
	NetworkUUID            *string `json:"network_uuid,omitempty"` // rubix-ce needs this
	PointUUID              *string `json:"point_uuid,omitempty"`
	WritePointUUID         *string `json:"write_point_uuid,omitempty"`
	ScheduleUUID           *string `json:"schedule_uuid,omitempty"`
	Query                  *string `json:"query,omitempty"`
}
