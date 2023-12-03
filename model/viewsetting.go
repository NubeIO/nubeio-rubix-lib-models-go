package model

import "gorm.io/datatypes"

type ViewSetting struct {
	CommonUUID
	Logo         datatypes.JSON `json:"logo"`
	Theme        datatypes.JSON `json:"theme"`
	WidgetConfig datatypes.JSON `json:"widget_config"`
}
