package model

type Plugin struct {
	Name        string
	PluginName  string
	Description string
}

var PluginsNames = struct {
	BACnetServer Plugin
	LoRa         Plugin
	Edge28       Plugin
}{
	BACnetServer: Plugin{Name: "BACnet-Server", PluginName: "bacnetserver", Description: "a plugin for an api to the nube bacnet-server"},
	LoRa:         Plugin{Name: "Nube LoRa Stack", PluginName: "lora", Description: "a network for nube lora sensors and IO controllers"},
	Edge28:       Plugin{Name: "Nube IO edge-28", PluginName: "edge28", Description: "a network for nube edge-28 controllers"},
}
