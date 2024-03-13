package dto

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"time"
)

type HistoryRequest struct {
	Filter        *string `json:"filter"`
	FilterHistory *string `json:"filter_history"`
}

type HistoryResponse struct {
	Data     []*HistoryData   `json:"data"`
	Tags     *HistoryTags     `json:"tags,omitempty"`
	MetaTags *HistoryMetaTags `json:"meta_tags,omitempty"`
}

type HistoryData struct {
	LocationUUID        string          `json:"location_uuid"`
	LocationName        string          `json:"location_name"`
	LocationDescription string          `json:"location_description"`
	GroupUUID           string          `json:"group_uuid"`
	GroupName           string          `json:"group_name"`
	GroupDescription    string          `json:"group_description"`
	HostUUID            string          `json:"host_uuid"`
	HostName            string          `json:"host_name"`
	HostDescription     string          `json:"host_description"`
	GlobalUUID          string          `json:"global_uuid"`
	NetworkUUID         string          `json:"network_uuid"`
	NetworkName         string          `json:"network_name"`
	NetworkDescription  string          `json:"network_description"`
	DeviceUUID          string          `json:"device_uuid"`
	DeviceName          string          `json:"device_name"`
	DeviceDescription   string          `json:"device_description"`
	PointUUID           string          `json:"point_uuid"`
	PointName           string          `json:"point_name"`
	PointDescription    string          `json:"point_description"`
	Values              []*HistoryValue `json:"values"`
}

type HistoryValue struct {
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type HistoryTags struct {
	Network map[string][]*HistoryTag `json:"network,omitempty"`
	Device  map[string][]*HistoryTag `json:"device,omitempty"`
	Point   map[string][]*HistoryTag `json:"point,omitempty"`
}

type HistoryMetaTags struct {
	Network map[string][]*HistoryMetaTag `json:"network,omitempty"`
	Device  map[string][]*HistoryMetaTag `json:"device,omitempty"`
	Point   map[string][]*HistoryMetaTag `json:"point,omitempty"`
}

type HistoryTag struct {
	Tag string `json:"tag"`
}

type HistoryMetaTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HistorySync struct {
	PointHistories []*model.PointHistory `json:"point_histories"`
	MetricLogs     []*model.MetricLog    `json:"metric_logs"`
}
