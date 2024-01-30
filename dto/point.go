package dto

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"time"
)

type PublishPointList struct {
	PluginName  string `json:"plugin_name"`
	NetworkName string `json:"network_name"`
	DeviceName  string `json:"device_name"`
	PointUUID   string `json:"point_uuid"`
	PointName   string `json:"point_name"`
}

// TimeOverride TODO add in later
// TimeOverride where a point value can be overridden for a duration of time
type TimeOverride struct {
	PointUUID string `json:"point_uuid" gorm:"references points;not null;default:null;primaryKey"`
	StartDate string `json:"start_date"` // START at 25:11:2021:13:00
	EndDate   string `json:"end_date"`   // START at 25:11:2021:13:30
	Value     string `json:"value"`
	Priority  string `json:"priority"`
}

// MathOperation same as in lora and point-server TODO add in later
type MathOperation struct {
	Calc string // x + 1
	X    float64
}

type PointWithParent struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	DeviceUUID  string `json:"device_uuid"`
	DeviceName  string `json:"device_name"`
	NetworkUUID string `json:"network_uuid"`
	NetworkName string `json:"network_name"`
}

type PointHistoryInterval struct {
	UUID            string   `json:"uuid"`
	HistoryInterval *int     `json:"history_interval,omitempty"`
	Timestamp       string   `json:"timestamp,omitempty"`
	PresentValue    *float64 `json:"present_value,omitempty"`
}

type PointForPostgresSync struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	DeviceUUID   string `json:"device_uuid,omitempty"`
	DeviceName   string `json:"device_name,omitempty"`
	NetworkUUID  string `json:"network_uuid"`
	NetworkName  string `json:"network_name"`
	GlobalUUID   string `json:"global_uuid"`
	HostUUID     string `json:"host_uuid"`
	HostName     string `json:"host_name"`
	GroupUUID    string `json:"group_uuid"`
	GroupName    string `json:"group_name"`
	LocationUUID string `json:"location_uuid"`
	LocationName string `json:"location_name"`
}

type PointTagForPostgresSync struct {
	HostUUID  string `json:"host_uuid"`
	PointUUID string `json:"point_uuid"`
	Tag       string `json:"tag"`
}

type PointMetaTagForPostgresSync struct {
	HostUUID  string `json:"host_uuid,omitempty"`
	PointUUID string `json:"point_uuid,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
}

type UpdatePointOpts struct {
	WriteValue bool
}

type PointWriter struct {
	Priority     *map[string]*float64 `json:"priority"`
	PresentValue *float64             `json:"present_value"`
	ForceWrite   bool                 `json:"force_write"`
	Message      string               `json:"message"`
	Fault        bool                 `json:"fault"`
	PollState    datatype.PointState  `json:"poll_state"`
	Timestamp    *time.Time           `json:"timestamp"`
}
type PointWriteResponse struct {
	Point                model.Point `json:"point"`
	IsPresentValueChange bool        `json:"is_present_value_change"`
	IsWriteValueChange   bool        `json:"is_write_value_change"`
	IsPriorityChanged    bool        `json:"is_priority_changed"`
}