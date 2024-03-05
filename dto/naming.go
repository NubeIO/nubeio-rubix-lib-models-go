package dto

var CommonNaming = struct {
	NubeIO                    string
	Plugin                    string
	Read                      string
	Write                     string
	Sync                      string
	Job                       string
	Schedule                  string
	Mapping                   string
	Rubix                     string
	RubixGlobal               string
	History                   string
	Histories                 string
	Node                      string
	TransportType             string
	Token                     string
	Location                  string
	Group                     string
	Host                      string
	HostComment               string
	SnapshotCreateLog         string
	SnapshotRestoreLog        string
	Member                    string
	MemberDevice              string
	View                      string
	ViewWidget                string
	ViewSetting               string
	ViewTemplate              string
	ViewTemplateWidget        string
	ViewTemplateWidgetPointer string
	Team                      string
	Ticket                    string
	TicketComment             string
	StreamLog                 string
	AlertTransaction          string
	Email                     string
	EmailAttachment           string
	AlertCondition            string
}{
	NubeIO:                    "Nube-IO",
	Plugin:                    "plugin",
	Read:                      "read",
	Write:                     "write",
	Sync:                      "sync",
	Job:                       "job",
	Schedule:                  "schedule",
	Mapping:                   "mapping",
	Rubix:                     "rubix",
	RubixGlobal:               "rubix_global",
	History:                   "history",
	Histories:                 "histories",
	Node:                      "node",
	TransportType:             "transport_type",
	Token:                     "token",
	Location:                  "location",
	Group:                     "group",
	Host:                      "host",
	HostComment:               "host_comment",
	SnapshotCreateLog:         "snapshot_create_log",
	SnapshotRestoreLog:        "shapshot_restore_log",
	Member:                    "member",
	MemberDevice:              "member_device",
	View:                      "view",
	ViewWidget:                "view_widget",
	ViewSetting:               "view_setting",
	ViewTemplate:              "view_template",
	ViewTemplateWidget:        "view_template_widget",
	ViewTemplateWidgetPointer: "view_template_widget_pointer",
	Team:                      "team",
	Ticket:                    "ticket",
	TicketComment:             "ticket_comment",
	StreamLog:                 "stream_log",
	AlertTransaction:          "alert_transaction",
	Email:                     "email",
	EmailAttachment:           "email_attachment",
	AlertCondition:            "alert_condition",
}

var ThingClass = struct {
	API            string `json:"api"`
	Network        string `json:"network"`
	Device         string `json:"device"`
	Point          string `json:"point"`
	InternalAPI    string `json:"internal_api"`
	ExternalAPI    string `json:"external_api"`
	Schedule       string `json:"schedule"`
	GlobalSchedule string `json:"global_schedule"`
	Alert          string `json:"alert"`
	Integration    string `json:"integration"`
	Message        string `json:"message"`
}{
	API:            "api",
	Network:        "network",
	Device:         "device",
	Point:          "point",
	InternalAPI:    "internal_api",
	ExternalAPI:    "external_api",
	Schedule:       "schedule",
	GlobalSchedule: "global_schedule",
	Alert:          "alert",
	Integration:    "integration",
	Message:        "message",
}

var ThingType = struct {
	API            string
	Network        string
	Device         string
	Point          string
	InternalAPI    string
	ExternalAPI    string
	Schedule       string
	GlobalSchedule string
	Alert          string
}{
	API:            "api",
	Network:        "network",
	Device:         "device",
	Point:          "point",
	InternalAPI:    "internal_api",
	ExternalAPI:    "external_api",
	Schedule:       "schedule",
	GlobalSchedule: "global_schedule",
	Alert:          "alert",
}

var WriterActions = struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Patch  string `json:"patch"`
	Delete string `json:"delete"`
}{
	Read:   "read",
	Write:  "write",
	Patch:  "patch",
	Delete: "delete",
}

var CommonFaultCode = struct {
	ConfigError      string
	SystemError      string
	PluginNotEnabled string
	Offline          string
	Ok               string
	PointError       string
	NetworkError     string
	DeviceError      string
	PointWriteOk     string
	PointWriteError  string
	LowBattery       string
}{
	ConfigError:      "configError",
	SystemError:      "systemError",
	PluginNotEnabled: "pluginNotEnabled",
	Offline:          "offline",
	Ok:               "ok",
	PointError:       "pointError",
	NetworkError:     "NetworkError",
	DeviceError:      "DeviceError",
	PointWriteOk:     "PointWriteOk",
	PointWriteError:  "PointWriteError",
	LowBattery:       "LowBattery",
}

var CommonFaultMessage = struct {
	ConfigError      string
	SystemError      string
	PluginNotEnabled string
	Offline          string
	NetworkMessage   string
}{
	ConfigError:      "config error",
	SystemError:      "system error",
	PluginNotEnabled: "plugin not enabled or no valid message from the network",
	Offline:          "offline",
	NetworkMessage:   "msg for network valid",
}

var MessageLevel = struct {
	Info         string
	Critical     string
	NoneCritical string
	Warning      string
	Fail         string
	Normal       string
}{
	Info:         "info",
	Critical:     "critical",
	NoneCritical: "noneCritical",
	Warning:      "warning",
	Fail:         "fail",
	Normal:       "normal",
}

var TransType = struct {
	Serial string `json:"serial"`
	IP     string `json:"ip"`
	LoRa   string `json:"LoRa"`
}{
	Serial: "serial",
	IP:     "ip",
	LoRa:   "LoRa",
}

var TransClient = struct {
	Client          string
	Server          string
	WirelessGateway string
	Stream          string
}{
	Client:          "client",
	Server:          "server",
	WirelessGateway: "wireless",
	Stream:          "gateway",
}

var TransProtocol = struct {
	REST         string
	BACnet       string
	Modbus       string
	ModbusMaster string
	MQTT         string
	LoRa         string
	Lora         string
	LoRaWAN      string
}{
	REST:         "rest",
	BACnet:       "BACnet",
	Modbus:       "Modbus",
	ModbusMaster: "ModbusMaster",
	MQTT:         "MQTT",
	LoRa:         "LoRa",
	Lora:         "lora",
	LoRaWAN:      "LoRaWAN",
}

var PointTags = struct {
	RSSI     string
	Voltage  string
	Temp     string
	Humidity string
	Light    string
	Motion   string
	Pulse    string
	AI1      string
	AI2      string
	AI3      string
}{
	RSSI:     "rssi",
	Voltage:  "voltage",
	Temp:     "temperature",
	Humidity: "humidity",
	Light:    "light",
	Motion:   "motion",
	Pulse:    "pulse",
	AI1:      "ai1",
	AI2:      "ai2",
	AI3:      "ai3",
}
