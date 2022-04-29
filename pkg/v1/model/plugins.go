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
	RubixIO      Plugin
	System       Plugin
}{
	BACnetServer: Plugin{Name: "BACnet-Server", PluginName: "bacnetserver", Description: "a plugin for an api to the nube bacnet-server"},
	LoRa:         Plugin{Name: "Nube LoRa Stack", PluginName: "lora", Description: "a network for nube lora sensors and IO controllers"},
	Edge28:       Plugin{Name: "Nube IO edge-28", PluginName: "edge28", Description: "a network for nube edge-28 controllers"},
	RubixIO:      Plugin{Name: "Nube IO rubix-io", PluginName: "rubix-io", Description: "a network for nube rubix-io controllers"},
	System:       Plugin{Name: "Nube IO system", PluginName: "system", Description: "a plugin for generic points and schedules"},
}
