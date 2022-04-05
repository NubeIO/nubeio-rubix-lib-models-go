package model

type PointMapping struct {
	AutoMappingNetworksSelection []string `json:"auto_mapping_networks_selection"` //user selected networks
	AutoMappingFlowNetworkName   string   `json:"auto_mapping_flow_network_name"`  //user can send name
	AutoMappingFlowNetworkUUID   string   `json:"auto_mapping_flow_network_uuid"`
	SourceNetworkUUID            string
	Point                        *Point `json:"point"`
}
