package model

type CommonDevice struct {
	Manufacture string  `json:"manufacture,omitempty"`                  // nube
	Model       string  `json:"model,omitempty"`                        // thml
	AddressId   int     `json:"address_id,omitempty"`                   // for example a modbus address or bacnet address
	ZeroMode    *bool   `json:"zero_mode,omitempty"`                    // if true means read from address 0 if false will read at 1
	AddressUUID *string `json:"address_uuid" gorm:"type:varchar(255);"` // AAB1213
	CommonIP
}

type Device struct {
	CommonUUID
	Name string `json:"name" gorm:"uniqueIndex:idx_devices_name_network_uuid"`
	CommonDescription
	CommonEnable
	CommonFault
	CommonCreated
	CommonThingClass
	CommonThingRef
	CommonThingType
	CommonDevice
	DeviceMac                *int             `json:"device_mac,omitempty"`
	DeviceObjectId           *int             `json:"device_object_id,omitempty"`
	NetworkNumber            *int             `json:"network_number,omitempty"` // bacnet network number
	MaxADPU                  *int             `json:"max_adpu,omitempty"`       // bacnet
	Segmentation             string           `json:"segmentation,omitempty"`   // bacnet
	DeviceMask               *int             `json:"device_mask,omitempty"`
	TypeSerial               *bool            `json:"type_serial,omitempty"`
	TransportType            string           `json:"transport_type,omitempty"` // serial, ip
	SupportsRpm              *bool            `json:"supports_rpm,omitempty"`   // bacnet support read property multiple
	SupportsWpm              *bool            `json:"supports_wpm,omitempty"`   // bacnet support write property multiple
	NetworkUUID              string           `json:"network_uuid,omitempty" gorm:"type:varchar(255) references networks;not null;default:null;uniqueIndex:idx_devices_name_network_uuid"`
	NumberOfDevicesPermitted *int             `json:"number_of_devices_permitted,omitempty"`
	Points                   []*Point         `json:"points,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tags                     []*Tag           `json:"tags,omitempty" gorm:"many2many:devices_tags;constraint:OnDelete:CASCADE"`
	FastPollRate             *float64         `json:"fast_poll_rate"`
	NormalPollRate           *float64         `json:"normal_poll_rate"`
	SlowPollRate             *float64         `json:"slow_poll_rate"`
	DelayBetweenPointsMs     *int             `json:"delay_between_points_ms,omitempty"` // used for polling when we need a time delay between each point
	DeviceTimeout            *int             `json:"device_timeout,omitempty"`
	MetaTags                 []*DeviceMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Connection               string           `json:"connection" gorm:"default:Connected"`
	ConnectionMessage        *string          `json:"connection_message" gorm:""`
	CommonSourceUUID
	CommonHistoryEnable
}

type DeviceMetaTag struct {
	DeviceUUID string `json:"device_uuid,omitempty" gorm:"type:varchar(255) references devices;not null;default:null;primaryKey"`
	Key        string `json:"key,omitempty" gorm:"primaryKey"`
	Value      string `json:"value,omitempty"`
}
