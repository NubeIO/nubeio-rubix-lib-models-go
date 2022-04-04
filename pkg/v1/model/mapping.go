package model

type PointMapping struct {
	FlowNetworkName string   `json:"flow_network_name"`
	PluginsList     []string `json:"plugins_list"`
	Point           *Point   `json:"point"`
}
