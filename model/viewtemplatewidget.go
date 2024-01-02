package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"gorm.io/datatypes"
)

type ViewTemplateWidget struct {
	CommonUUID
	ViewTemplateUUID           string                       `json:"view_template_uuid,omitempty" gorm:"type:varchar(255) references view_templates;"`
	Name                       string                       `json:"name" gorm:"type:varchar(255);not null;"`
	Order                      int                          `json:"order"`
	X                          int                          `json:"x"`
	Y                          int                          `json:"y"`
	Type                       *string                      `json:"type"`
	Config                     datatypes.JSON               `json:"config"`
	Class                      datatype.Class               `json:"class,omitempty"`
	HasDiffRW                  *bool                        `json:"has_diff_rw,omitempty"`
	NetworkName                *string                      `json:"network_name,omitempty"`
	DeviceName                 *string                      `json:"device_name,omitempty"`
	PointName                  *string                      `json:"point_name,omitempty"`
	WriteNetworkName           *string                      `json:"write_network_name,omitempty"`
	WriteDeviceName            *string                      `json:"write_device_name,omitempty"`
	WritePointName             *string                      `json:"write_point_name,omitempty"`
	ScheduleName               *string                      `json:"schedule_name,omitempty"`
	Query                      *string                      `json:"query,omitempty"`
	ViewTemplateWidgetPointers []*ViewTemplateWidgetPointer `json:"view_template_widget_pointer,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}
