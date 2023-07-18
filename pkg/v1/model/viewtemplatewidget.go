package model

import "gorm.io/datatypes"

type ViewTemplateWidget struct {
	CommonUUID
	ViewTemplateUUID           string                       `json:"view_template_uuid,omitempty" gorm:"type:varchar(255) references view_templates;"`
	Name                       string                       `json:"name" gorm:"type:varchar(255);not null;"`
	Order                      int                          `json:"order"`
	X                          int                          `json:"x"`
	Y                          int                          `json:"y"`
	Type                       *string                      `json:"type"`
	Config                     datatypes.JSON               `json:"config"`
	Class                      string                       `json:"class,omitempty"`
	NetworkName                *string                      `json:"network_name,omitempty"`
	DeviceName                 *string                      `json:"device_name,omitempty"`
	PointName                  *string                      `json:"point_name,omitempty"`
	ScheduleName               *string                      `json:"schedule_name,omitempty"`
	ViewTemplateWidgetPointers []*ViewTemplateWidgetPointer `json:"view_template_widget_pointer" gorm:"constraint:OnDelete:CASCADE;"`
}
