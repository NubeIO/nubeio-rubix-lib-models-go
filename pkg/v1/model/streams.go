package model

type CommonStreamAutoMapping struct {
	AutoMappingNetworkUUID  *string `json:"auto_mapping_network_uuid"`
	AutoMappingDeviceUUID   *string `json:"auto_mapping_device_uuid"`
	AutoMappingScheduleUUID *string `json:"auto_mapping_schedule_uuid"`
}

type CommonStream struct {
	CommonUUID
	CommonSyncUUID
	CommonDescription
	CommonEnable
	CommonCreated
	CommonStreamAutoMapping
}

type Stream struct {
	CommonNameUnique
	CommonStream
	FlowNetworks  []*FlowNetwork  `json:"flow_networks" gorm:"many2many:flow_networks_streams;constraint:OnDelete:CASCADE"`
	Producers     []*Producer     `json:"producers" gorm:"constraint:OnDelete:CASCADE"`
	CommandGroups []*CommandGroup `json:"command_groups" gorm:"constraint:OnDelete:CASCADE"`
	Tags          []*Tag          `json:"tags" gorm:"many2many:streams_tags;constraint:OnDelete:CASCADE"`
	CommonCreatedFromAutoMapping
}

type StreamClone struct {
	Name string `json:"name" gorm:"uniqueIndex:idx_stream_clones_name_flow_network_clone_uuid"`
	CommonStream
	CommonSourceUUIDUnique
	CommonConnection
	FlowNetworkCloneUUID string      `json:"flow_network_clone_uuid" gorm:"type:string references flow_network_clones;not null;default:null;uniqueIndex:idx_stream_clones_name_flow_network_clone_uuid"`
	Consumers            []*Consumer `json:"consumers" gorm:"constraint:OnDelete:CASCADE"`
	Tags                 []*Tag      `json:"tags" gorm:"many2many:stream_clones_tags;constraint:OnDelete:CASCADE"`
	CommonCreatedFromAutoMapping
}

type SyncStream struct {
	GlobalUUID string
	Stream     *Stream
}
