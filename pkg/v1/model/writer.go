package model

import "gorm.io/datatypes"

type CommonWriter struct {
	CommonUUID
	CommonSyncUUID
	WriterThingClass string         `json:"writer_thing_class,omitempty"`
	WriterThingType  string         `json:"writer_thing_type,omitempty"`
	WriterThingUUID  string         `json:"writer_thing_uuid,omitempty"`
	DataStore        datatypes.JSON `json:"data_store,omitempty"`
	CommonConnection
	CommonCreated
	CommonCreatedFromAutoMapping
}

type Writer struct {
	CommonWriter
	WriterThingName string   `json:"writer_thing_name,omitempty" gorm:"uniqueIndex:idx_writers_consumer_uuid_writer_thing_name"`
	ConsumerUUID    string   `json:"consumer_uuid,omitempty" gorm:"type:string references consumers;not null;default:null;uniqueIndex:idx_writers_consumer_uuid_writer_thing_name"`
	PresentValue    *float64 `json:"present_value,omitempty"`
}

type WriterClone struct { // TODO the WriterClone needs to publish a COV event as for example we have 2x mqtt broker then the cov for a point maybe different when not going over the internet
	CommonWriter
	WriterThingName   string `json:"writer_thing_name,omitempty"`
	ProducerUUID      string `json:"producer_uuid" gorm:"type:string references producers;not null;default:null"`
	FlowFrameworkUUID string `json:"flow_framework_uuid,omitempty"`
	CommonSourceUUIDUnique
}

type SyncWriter struct {
	ProducerUUID      string
	WriterUUID        string
	FlowFrameworkUUID string
}

type WriterBody struct {
	Action   string               `json:"action,omitempty"`
	Priority *map[string]*float64 `json:"priority,omitempty"`
	Schedule *ScheduleData        `json:"schedule,omitempty"`
}

type WriterBulkBody struct {
	WriterUUID string `json:"writer_uuid,omitempty"`
	WriterBody
}

type WriterActionOutput struct {
	UUID         string          `json:"uuid"`
	Action       string          `json:"action"`
	IsError      bool            `json:"is_error"`
	Message      *string         `json:"message,omitempty"`
	PresentValue *float64        `json:"present_value,omitempty"`
	DataStore    *datatypes.JSON `json:"data_store,omitempty"`
}

type SyncCOV struct {
	Priority     *map[string]*float64
	PresentValue *float64
	Schedule     *ScheduleData
}

type SyncWriterAction struct {
	Priority *map[string]*float64
	Schedule *ScheduleData
}

type PointWriter struct {
	Priority     *map[string]*float64
	PresentValue *float64
}
