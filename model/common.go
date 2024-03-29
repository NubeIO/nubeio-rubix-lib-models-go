package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"time"
)

type CommonDescription struct {
	Description string `json:"description,omitempty"`
}

type CommonName struct {
	Name string `json:"name"`
}

type CommonThingClass struct {
	ThingClass string `json:"thing_class,omitempty"`
}

type CommonThingType struct {
	ThingType string `json:"thing_type,omitempty"`
}

type CommonThingObject struct {
	ThingObject string `json:"thing_object,omitempty"`
}

type CommonThingRef struct {
	ThingRef string `json:"thing_reference,omitempty"`
}

type CommonThingUUID struct {
	ThingUUID string `json:"thing_uuid,omitempty"`
}

type CommonNameUnique struct {
	Name string `json:"name"  gorm:"type:varchar(255);unique;not null"`
}

type CommonModulePath struct {
	ModulePath string `json:"module_path"  gorm:"type:varchar(255);unique;not null"`
}

type CommonHelp struct {
	Help string `json:"help"`
}

type CommonIP struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

type CommonNodeType struct {
	NodeType string `json:"node_type"`
}

type CommonType struct {
	ObjectType string `json:"object_type"`
}

type CommonAction struct {
	Action string `json:"action"`
}

type CommonEnable struct {
	Enable *bool `json:"enable" gorm:"default:false"`
}

type CommonID struct {
	ID string `json:"id"`
}

type CommonIDUnique struct {
	Name string `json:"id" gorm:"type:varchar(255);unique;not null"`
}

type CommonUUID struct {
	UUID string `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
}

type CommonSourceUUIDUnique struct {
	SourceUUID string `json:"source_uuid" gorm:"uniqueIndex;not null"`
}

type CommonSourceUUID struct {
	SourceUUID *string `json:"source_uuid"`
}

type CommonSyncUUID struct {
	SyncUUID string `json:"sync_uuid"`
}

type CommonConnection struct {
	Connection string `json:"connection" gorm:"default:Connected"`
	Message    string `json:"message" gorm:"default:N/A"`
}

type CommonCreated struct {
	CreatedAt time.Time `json:"created_on,omitempty"`
	UpdatedAt time.Time `json:"updated_on,omitempty"` // updated from a plugin or from UI
	LastWrite time.Time `json:"last_write,omitempty"` // timestamp for last time user wrote a value from CE
}

type CommonHistoryEnable struct {
	HistoryEnable *bool `json:"history_enable" gorm:"default:false"`
}

type CommonFault struct {
	InFault      bool      `json:"fault,omitempty"`
	MessageLevel string    `json:"message_level,omitempty"`
	MessageCode  string    `json:"message_code,omitempty"`
	Message      string    `json:"message,omitempty"`
	MessageFail  string    `json:"message_fail,omitempty"`
	LastOk       time.Time `json:"last_ok,omitempty"`
	LastFail     time.Time `json:"last_fail,omitempty"`
}

type CommonDevice struct {
	Manufacture string  `json:"manufacture,omitempty"`                  // nube
	Model       string  `json:"model,omitempty"`                        // thml
	AddressId   int     `json:"address_id,omitempty"`                   // for example a modbus address or bacnet address
	ZeroMode    *bool   `json:"zero_mode,omitempty"`                    // if true means read from address 0 if false will read at 1
	AddressUUID *string `json:"address_uuid" gorm:"type:varchar(255);"` // AAB1213
	CommonIP
}

// HistoryConfig for configuring point histories
type HistoryConfig struct {
	CommonHistoryEnable
	HistoryType         datatype.HistoryType `json:"history_type,omitempty"`
	HistoryInterval     *int                 `json:"history_interval,omitempty"`
	HistoryCOVThreshold *float64             `json:"history_cov_threshold,omitempty"`
}
