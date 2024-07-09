package model

type Host struct {
	UUID                 string  `json:"uuid" gorm:"primaryKey"`
	GlobalUUID           string  `json:"global_uuid"`
	DeviceType           string  `json:"device_type"`
	Name                 string  `json:"name"  gorm:"type:varchar(255);not null;uniqueIndex:idx_hosts_name_group_uuid"`
	GroupUUID            string  `json:"group_uuid,omitempty" gorm:"TYPE:varchar(255) REFERENCES groups;not null;default:null;uniqueIndex:idx_hosts_name_group_uuid"`
	Enable               *bool   `json:"enable"`
	Description          string  `json:"description"`
	IP                   string  `json:"ip"`
	BiosPort             int     `json:"bios_port"`
	Port                 int     `json:"port"`
	HTTPS                *bool   `json:"https"`
	IsOnline             *bool   `json:"is_online"`
	IsValidToken         *bool   `json:"is_valid_token"`
	PingFailCount        int     `json:"ping_fail_count"`
	ExternalToken        string  `json:"external_token"`
	VirtualIP            string  `json:"virtual_ip"`
	ReceivedBytes        int     `json:"received_bytes"`
	SentBytes            int     `json:"sent_bytes"`
	ConnectedSince       string  `json:"connected_since"`
	Timezone             string  `json:"timezone"`
	ROSVersion           string  `json:"ros_version"`
	ROSRestartExpression *string `json:"ros_restart_expression,omitempty"`
	RebootExpression     *string `json:"reboot_expression,omitempty"`
	CommonHistoryEnable
	Tags     []*Tag         `json:"tags,omitempty" gorm:"many2many:hosts_tags;constraint:OnDelete:CASCADE"`
	Comments []*HostComment `json:"comments,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Views    []*View        `json:"views,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Networks []*Network     `json:"networks,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
