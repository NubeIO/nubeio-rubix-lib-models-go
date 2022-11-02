package model

import (
	"gorm.io/datatypes"
	"time"
)

// ProducerHistory for storing the history
type ProducerHistory struct {
	ID           int    `json:"id" gorm:"autoIncrement;primaryKey;index"`
	ProducerUUID string `json:"producer_uuid,omitempty" gorm:"type:varchar(255) references producers;not null;default:null"`
	CommonCurrentWriterUUID
	PresentValue *float64        `json:"present_value"`
	DataStore    *datatypes.JSON `json:"data_store,omitempty"`
	Timestamp    time.Time       `json:"timestamp,omitempty"`
}
