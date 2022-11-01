package model

type CommonStream struct {
	CommonUUID
	CommonSyncUUID
	CommonDescription
	CommonEnable
	CommonCreated
}

type Stream struct {
	CommonNameUnique
	CommonStream
	FlowNetworks  []*FlowNetwork  `json:"flow_networks" gorm:"many2many:flow_networks_streams;constraint:OnDelete:CASCADE"`
	Producers     []*Producer     `json:"producers" gorm:"constraint:OnDelete:CASCADE;"`
	CommandGroups []*CommandGroup `json:"command_groups" gorm:"constraint:OnDelete:CASCADE;"`
	Tags          []*Tag          `json:"tags" gorm:"many2many:streams_tags;constraint:OnDelete:CASCADE"`
}

type StreamClone struct {
	Name string `json:"name" gorm:"uniqueIndex:idx_stream_clones_name_flow_network_clone_uuid"`
	CommonStream
	CommonSourceUUID
	CommonConnection
	FlowNetworkCloneUUID string      `json:"flow_network_clone_uuid" gorm:"TYPE:string REFERENCES flow_network_clones;not null;default:null;uniqueIndex:idx_stream_clones_name_flow_network_clone_uuid"`
	Consumers            []*Consumer `json:"consumers" gorm:"constraint:OnDelete:CASCADE;"`
	Tags                 []*Tag      `json:"tags" gorm:"many2many:stream_clones_tags;constraint:OnDelete:CASCADE"`
}

type SyncStream struct {
	GlobalUUID string
	Stream     *Stream
}
