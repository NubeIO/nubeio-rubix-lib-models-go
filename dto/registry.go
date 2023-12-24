package dto

type DeviceInfo struct {
	GlobalUUID string  `json:"global_uuid"`
	Version    string  `json:"version"`
	Type       string  `json:"type"`
	ROS        ROSInfo `json:"ros"`
}

type ProductInfo struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}

type DeviceInfoFirstRecord struct { // TODO: remove after migration done
	DeviceInfo DeviceInfo `json:"1"`
}

type DeviceInfoDefault struct { // TODO: remove after migration done
	DeviceInfoFirstRecord DeviceInfoFirstRecord `json:"_default"`
}

type ROSInfo struct {
	Version           string  `json:"version"`
	RestartExpression *string `json:"restart_expression,omitempty"`
}
