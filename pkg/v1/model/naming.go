package model

var CommonNaming = struct {
	NubeIO           string
	Plugin           string
	Read             string
	Write            string
	Sync             string
	Stream           string
	StreamClone      string
	StreamList       string
	Job              string
	Schedule         string
	Producer         string
	Consumer         string
	Writer           string
	WriterClone      string
	Mapping          string
	CommandGroup     string
	Rubix            string
	RubixGlobal      string
	FlowNetwork      string
	FlowNetworkClone string
	History          string
	ProducerHistory  string
	Histories        string
	Node             string
	TransportType    string
}{
	NubeIO:           "Nube-IO",
	Plugin:           "plugin",
	Read:             "read",
	Write:            "write",
	Sync:             "sync",
	Stream:           "stream",
	StreamClone:      "stream_clone",
	StreamList:       "stream_list",
	Job:              "job",
	Schedule:         "schedule",
	Producer:         "producer",
	Consumer:         "consumer",
	Writer:           "writer",
	WriterClone:      "writer_clone",
	Mapping:          "mapping",
	CommandGroup:     "command_group",
	Rubix:            "rubix",
	RubixGlobal:      "rubix_global",
	FlowNetwork:      "flow_network",
	FlowNetworkClone: "flow_network_clone",
	History:          "history",
	ProducerHistory:  "producer_history",
	Histories:        "histories",
	Node:             "node",
	TransportType:    "transport_type",
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
}{
	ConfigError:      "configError",
	SystemError:      "systemError",
	PluginNotEnabled: "pluginNotEnabled",
	Offline:          "offline",
	Ok:               "ok",
	PointError:       "pointError",
	NetworkError:     "NetworkError",
	DeviceError:      "DeviceError",
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
