package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"gorm.io/datatypes"
	"reflect"
	"time"
)

// TimeOverride TODO add in later
// TimeOverride where a point value can be overridden for a duration of time
type TimeOverride struct {
	PointUUID string `json:"point_uuid" gorm:"references points;not null;default:null;primaryKey"`
	StartDate string `json:"start_date"` // START at 25:11:2021:13:00
	EndDate   string `json:"end_date"`   // START at 25:11:2021:13:30
	Value     string `json:"value"`
	Priority  string `json:"priority"`
}

// MathOperation same as in lora and point-server TODO add in later
type MathOperation struct {
	Calc string // x + 1
	X    float64
}

// Point table
type Point struct {
	CommonUUID
	Name string `json:"name" gorm:"uniqueIndex:idx_points_name_device_uuid"`
	CommonDescription
	CommonEnable
	CommonCreated
	CommonThingClass
	CommonThingRef
	CommonThingType
	CommonFault
	PresentValue           *float64                        `json:"present_value"` // point value, read only
	OriginalValue          *float64                        `json:"original_value"`
	DisplayValue           *string                         `json:"display_value"`        // 248.67 °F
	WriteValue             *float64                        `json:"write_value"`          // writeValue was added so if user wanted to do a math function on the point write
	WriteValueOriginal     *float64                        `json:"write_value_original"` // writeValue was added so if user wanted to do a math function on the point write
	CurrentPriority        *int                            `json:"current_priority,omitempty"`
	WritePriority          *int                            `json:"write_priority,omitempty"` // used for user just to select the write priority on example for bacnet
	IsOutput               *bool                           `json:"is_output"`                // used for as example: for bacnet-server we only support AV so if a point IsOutput = false then for mapping its set as a consumer but if true then its set as a prodcuer
	IsTypeBool             *bool                           `json:"is_type_bool"`
	InSync                 *bool                           `json:"in_sync"` // is set to false when a new value is written from the user example: if its false then modbus would write the new value. if user edits the point it will disable the COV for one time
	Fallback               *float64                        `json:"fallback"`
	DeviceUUID             string                          `json:"device_uuid,omitempty" gorm:"type:string references devices;not null;default:null;uniqueIndex:idx_points_name_device_uuid"`
	EnableWriteable        *bool                           `json:"writeable,omitempty"`             // TODO: UI should hide the `write` action if not enabled
	MathOnPresentValue     string                          `json:"math_on_present_value,omitempty"` // TODO: THIS SHOULD BE DELETED WHEN SAFE TO DO SO.
	MathOnWriteValue       string                          `json:"math_on_write_value,omitempty"`   // TODO: THIS SHOULD BE DELETED WHEN SAFE TO DO SO.
	COV                    *float64                        `json:"cov"`
	ObjectType             string                          `json:"object_type,omitempty"` // binaryInput, coil, if type os input don't return the priority array
	ObjectId               *int                            `json:"object_id,omitempty"`
	DataType               string                          `json:"data_type,omitempty"`       // int16, uint16, float32
	IsBitwise              *bool                           `json:"is_bitwise,omitempty"`      // true indicates data type is numeric, but we are taking only a specific bit of the number. Used for Modbus and Bacnet.
	BitwiseIndex           *int                            `json:"bitwise_index,omitempty"`   // if IsBitwise is set, then BitwiseIndex sets the index to read from numeric value
	ObjectEncoding         string                          `json:"object_encoding,omitempty"` // BEB_LEW bebLew
	IoNumber               string                          `json:"io_number,omitempty"`       // DI1,UI1,AO1, temp, pulse, motion
	IoType                 string                          `json:"io_type,omitempty"`         // 0-10dc, 0-40ma, thermistor
	AddressID              *int                            `json:"address_id"`                // for example a modbus address or bacnet address
	AddressLength          *int                            `json:"address_length"`            // for example a modbus address offset
	AddressUUID            *string                         `json:"address_uuid,omitempty"`    // for example a droplet id (so a string)
	NextAvailableAddress   *bool                           `json:"use_next_available_address,omitempty"`
	Decimal                *uint32                         `json:"decimal,omitempty"`
	MultiplicationFactor   *float64                        `json:"multiplication_factor"`
	ScaleEnable            *bool                           `json:"scale_enable"`
	ScaleInMin             *float64                        `json:"scale_in_min"`
	ScaleInMax             *float64                        `json:"scale_in_max"`
	ScaleOutMin            *float64                        `json:"scale_out_min"`
	ScaleOutMax            *float64                        `json:"scale_out_max"`
	Offset                 *float64                        `json:"offset"`
	UnitType               *string                         `json:"unit_type,omitempty"` // temperature
	Unit                   *string                         `json:"unit,omitempty"`
	UnitTo                 *string                         `json:"unit_to,omitempty"` // with take the unit and convert to, this would affect the presentValue and the original value will be stored in the raw
	Priority               *Priority                       `json:"priority,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tags                   []*Tag                          `json:"tags,omitempty" gorm:"many2many:points_tags;constraint:OnDelete:CASCADE"`
	ValueUpdatedFlag       *bool                           `json:"value_updated_flag,omitempty"`      // This is used when a plugin updates the PresentValue (not from priority array) and it triggers UpdatePointValue() to broadcast to producers. Should only be set to FALSE from UpdatePointValue().
	PointPriorityArrayMode datatype.PointPriorityArrayMode `json:"point_priority_use_type,omitempty"` // This configures how the point handles the priority array and present value.
	WriteMode              datatype.WriteMode              `json:"write_mode,omitempty"`
	ReadWriteType          datatype.ReadWriteType          `json:"read_write_type,omitempty"`
	WritePollRequired      *bool                           `json:"write_required,omitempty"`
	ReadPollRequired       *bool                           `json:"read_required,omitempty"`
	PollPriority           datatype.PollPriority           `json:"poll_priority"`
	PollRate               datatype.PollRate               `json:"poll_rate"`
	PointState             datatype.PointState             `json:"point_state,omitempty"`
	BACnetWriteToPV        *bool                           `json:"bacnet_write_to_pv,omitempty"`
	HistoryConfig
	MetaTags          []*PointMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Connection        string          `json:"connection" gorm:"default:Connected"`
	ConnectionMessage *string         `json:"connection_message" gorm:""`
	PointHistories    []*PointHistory `json:"point_histories,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	CommonSourceUUID
	LastHistoryValue     *float64       `json:"last_history_value,omitempty"`
	LastHistoryTimestamp *time.Time     `json:"last_history_timestamp,omitempty"`
	Config               datatypes.JSON `json:"config"`
}

type PointMetaTag struct {
	PointUUID string `json:"point_uuid,omitempty" gorm:"type:varchar(255) references points;not null;default:null;primaryKey"`
	Key       string `json:"key,omitempty" gorm:"primaryKey"`
	Value     string `json:"value,omitempty"`
}

type Priorities struct {
	P1  *float64 `json:"_1"`
	P2  *float64 `json:"_2"`
	P3  *float64 `json:"_3"`
	P4  *float64 `json:"_4"`
	P5  *float64 `json:"_5"`
	P6  *float64 `json:"_6"`
	P7  *float64 `json:"_7"`
	P8  *float64 `json:"_8"`
	P9  *float64 `json:"_9"`
	P10 *float64 `json:"_10"`
	P11 *float64 `json:"_11"`
	P12 *float64 `json:"_12"`
	P13 *float64 `json:"_13"`
	P14 *float64 `json:"_14"`
	P15 *float64 `json:"_15"`
	P16 *float64 `json:"_16"`
}

type Priority struct {
	PointUUID string   `json:"point_uuid,omitempty" gorm:"references points;not null;default:null;primaryKey"`
	P1        *float64 `json:"_1"`
	P2        *float64 `json:"_2"`
	P3        *float64 `json:"_3"`
	P4        *float64 `json:"_4"`
	P5        *float64 `json:"_5"`
	P6        *float64 `json:"_6"`
	P7        *float64 `json:"_7"`
	P8        *float64 `json:"_8"`
	P9        *float64 `json:"_9"`
	P10       *float64 `json:"_10"`
	P11       *float64 `json:"_11"`
	P12       *float64 `json:"_12"`
	P13       *float64 `json:"_13"`
	P14       *float64 `json:"_14"`
	P15       *float64 `json:"_15"`
	P16       *float64 `json:"_16"`
}

type PointWriter struct {
	Priority     *map[string]*float64 `json:"priority"`
	PresentValue *float64             `json:"present_value"`
	ForceWrite   bool                 `json:"force_write"`
	Message      string               `json:"message"`
	Fault        bool                 `json:"fault"`
	PollState    datatype.PointState  `json:"poll_state"`
	Timestamp    *time.Time           `json:"timestamp"`
}

func (p *Priority) GetHighestPriorityValue() *float64 {
	if p.P1 != nil {
		return p.P1
	}
	if p.P2 != nil {
		return p.P2
	}
	if p.P3 != nil {
		return p.P3
	}
	if p.P4 != nil {
		return p.P4
	}
	if p.P5 != nil {
		return p.P5
	}
	if p.P6 != nil {
		return p.P6
	}
	if p.P7 != nil {
		return p.P7
	}
	if p.P8 != nil {
		return p.P8
	}
	if p.P9 != nil {
		return p.P9
	}
	if p.P10 != nil {
		return p.P10
	}
	if p.P11 != nil {
		return p.P11
	}
	if p.P12 != nil {
		return p.P12
	}
	if p.P13 != nil {
		return p.P13
	}
	if p.P14 != nil {
		return p.P14
	}
	if p.P15 != nil {
		return p.P15
	}
	if p.P16 != nil {
		return p.P16
	}
	return nil
}

// GetPntType find the type of point as in voltage_dc
func GetPntType(strut interface{}, ioType string) (out string, err error) {
	val := reflect.ValueOf(strut).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.Name == ioType {
			return valueField.String(), nil
		}
	}
	return
}

type PointWithParent struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	DeviceUUID  string `json:"device_uuid"`
	DeviceName  string `json:"device_name"`
	NetworkUUID string `json:"network_uuid"`
	NetworkName string `json:"network_name"`
}

type PointHistoryInterval struct {
	UUID            string   `json:"uuid"`
	HistoryInterval *int     `json:"history_interval,omitempty"`
	Timestamp       string   `json:"timestamp,omitempty"`
	PresentValue    *float64 `json:"present_value,omitempty"`
}

type PointForPostgresSync struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	DeviceUUID   string `json:"device_uuid,omitempty"`
	DeviceName   string `json:"device_name,omitempty"`
	NetworkUUID  string `json:"network_uuid"`
	NetworkName  string `json:"network_name"`
	GlobalUUID   string `json:"global_uuid"`
	HostUUID     string `json:"host_uuid"`
	HostName     string `json:"host_name"`
	GroupUUID    string `json:"group_uuid"`
	GroupName    string `json:"group_name"`
	LocationUUID string `json:"location_uuid"`
	LocationName string `json:"location_name"`
}

type PointTagForPostgresSync struct {
	HostUUID  string `json:"host_uuid"`
	PointUUID string `json:"point_uuid"`
	Tag       string `json:"tag"`
}

type PointMetaTagForPostgresSync struct {
	HostUUID  string `json:"host_uuid,omitempty"`
	PointUUID string `json:"point_uuid,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
}

type UpdatePointOpts struct {
	WriteValue bool
}

type PointWriteResponse struct {
	Point                Point `json:"point"`
	IsPresentValueChange bool  `json:"is_present_value_change"`
	IsWriteValueChange   bool  `json:"is_write_value_change"`
	IsPriorityChanged    bool  `json:"is_priority_changed"`
}
