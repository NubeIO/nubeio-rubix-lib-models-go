package dto

type PublishPointList struct {
	PluginName  string `json:"plugin_name"`
	NetworkName string `json:"network_name"`
	DeviceName  string `json:"device_name"`
	PointUUID   string `json:"point_uuid"`
	PointName   string `json:"point_name"`
}
