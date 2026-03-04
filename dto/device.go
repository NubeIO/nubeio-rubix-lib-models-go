package dto

type DeviceTagForPostgresSync struct {
	HostUUID   string `json:"host_uuid"`
	DeviceUUID string `json:"device_uuid"`
	Tag        string `json:"tag"`
}

type DeviceMetaTagForPostgresSync struct {
	HostUUID   string `json:"host_uuid,omitempty"`
	DeviceUUID string `json:"device_uuid,omitempty"`
	Key        string `json:"key,omitempty"`
	Value      string `json:"value,omitempty"`
}

type SwapDevicePosition struct {
	UUID1 string `json:"uuid1"`
	UUID2 string `json:"uuid2"`
}
