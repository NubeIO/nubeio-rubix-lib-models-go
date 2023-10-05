package model

import (
	"reflect"
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

type ObjectType string

const (
	// bacnet
	ObjTypeAnalogInput  ObjectType = "analog_input"
	ObjTypeAnalogOutput ObjectType = "analog_output"
	ObjTypeAnalogValue  ObjectType = "analog_value"
	ObjTypeBinaryInput  ObjectType = "binary_input"
	ObjTypeBinaryOutput ObjectType = "binary_output"
	ObjTypeBinaryValue  ObjectType = "binary_value"
	ObjTypeEnumInput    ObjectType = "multi_state_input"
	ObjTypeEnumOutput   ObjectType = "multi_state_output"
	ObjTypeEnumValue    ObjectType = "multi_state_value"
	ObjAnalogInput      ObjectType = "analogInput"
	ObjAnalogOutput     ObjectType = "analogOutput"
	ObjAnalogValue      ObjectType = "analogValue"
	ObjBinaryInput      ObjectType = "binaryInput"
	ObjBinaryOutput     ObjectType = "binaryOutput"
	ObjBinaryValue      ObjectType = "binaryValue"
	ObjEnumInput        ObjectType = "multiStateInput"
	ObjEnumOutput       ObjectType = "multiStateOutput"
	ObjEnumValue        ObjectType = "multiStateValue"

	// modbus
	ObjTypeReadCoil           ObjectType = "read_coil"
	ObjTypeReadCoils          ObjectType = "read_coils"
	ObjTypeReadDiscreteInput  ObjectType = "read_discrete_input"
	ObjTypeReadDiscreteInputs ObjectType = "read_discrete_inputs"
	ObjTypeWriteCoil          ObjectType = "write_coil"
	ObjTypeWriteCoils         ObjectType = "write_coils"
	ObjTypeReadRegister       ObjectType = "read_register"
	ObjTypeReadRegisters      ObjectType = "read_registers"
	ObjTypeReadHolding        ObjectType = "read_holding"
	ObjTypeReadHoldings       ObjectType = "read_holdings"
	ObjTypeWriteHolding       ObjectType = "write_holding"
	ObjTypeWriteHoldings      ObjectType = "write_holdings"
	ObjTypeReadInt16          ObjectType = "read_int_16"
	ObjTypeWriteInt16         ObjectType = "write_int_16"
	ObjTypeReadUint16         ObjectType = "read_uint_16"
	ObjTypeWriteUint16        ObjectType = "write_uint_16"
	ObjTypeReadInt32          ObjectType = "read_int_32"
	ObjTypeReadUint32         ObjectType = "read_uint_32"
	ObjTypeReadFloat32        ObjectType = "read_float_32"
	ObjTypeWriteFloat32       ObjectType = "write_float_32"
	ObjTypeReadFloat64        ObjectType = "read_float_64"
	ObjTypeWriteFloat64       ObjectType = "write_float_64"
	ObjTypeCoil               ObjectType = "coil"
	ObjTypeDiscreteInput      ObjectType = "discrete_input"
	ObjTypeInputRegister      ObjectType = "input_register"
	ObjTypeHoldingRegister    ObjectType = "holding_register"
)

var ObjectTypesMap = map[ObjectType]int8{
	// bacnet
	ObjTypeAnalogInput: 0, ObjTypeAnalogOutput: 0, ObjTypeAnalogValue: 0,
	ObjTypeBinaryInput: 0, ObjTypeBinaryOutput: 0, ObjTypeBinaryValue: 0,
	ObjAnalogInput: 0, ObjAnalogOutput: 0, ObjAnalogValue: 0,
	ObjBinaryInput: 0, ObjBinaryOutput: 0, ObjBinaryValue: 0,
	ObjTypeEnumInput: 0, ObjTypeEnumOutput: 0, ObjTypeEnumValue: 0,
	// modbus
	ObjTypeReadCoil: 0, ObjTypeReadCoils: 0, ObjTypeReadDiscreteInput: 0,
	ObjTypeReadDiscreteInputs: 0, ObjTypeWriteCoil: 0, ObjTypeWriteCoils: 0,
	ObjTypeReadRegister: 0, ObjTypeReadRegisters: 0, ObjTypeReadHolding: 0,
	ObjTypeReadHoldings: 0, ObjTypeWriteHolding: 0, ObjTypeWriteHoldings: 0,
	ObjTypeReadInt16: 0, ObjTypeWriteInt16: 0,
	ObjTypeReadUint16: 0, ObjTypeWriteUint16: 0,
	ObjTypeReadInt32:   0,
	ObjTypeReadUint32:  0,
	ObjTypeReadFloat32: 0, ObjTypeWriteFloat32: 0,
	ObjTypeReadFloat64: 0, ObjTypeWriteFloat64: 0,
	ObjTypeCoil: 0, ObjTypeDiscreteInput: 0, ObjTypeInputRegister: 0, ObjTypeHoldingRegister: 0,
}

type DataType string

const (
	TypeDigital   DataType = "digital"
	TypeInt16     DataType = "int16"
	TypeUint16    DataType = "uint16"
	TypeInt32     DataType = "int32"
	TypeUint32    DataType = "uint32"
	TypeInt64     DataType = "int64"
	TypeUint64    DataType = "uint64"
	TypeFloat32   DataType = "float32"
	TypeFloat64   DataType = "float64"
	TypeMod10_u32 DataType = "mod10-u32" // MOD10k unsigned 32bit (2 registers) - A 32-bit unsigned value returned in two consecutive 16-bit registers. R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10_s32 DataType = "mod10-s32" // TODO: (not yet implemented) - MOD10k signed 32bit (2 registers) - A 32-bit signed value returned in two consecutive 16-bit registers. R2*10,000 + R1.  Each registers range is -9,999 to +9,999
	TypeMod10_u48 DataType = "mod10-u48" // TODO: (not yet implemented) - MOD10k unsigned 48bit (2 registers) - A 48-bit unsigned value returned in three consecutive 16-bit registers. R3*10,000^2 + R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10_s48 DataType = "mod10-s48" // TODO: (not yet implemented) - MOD10k signed 48bit (2 registers) - A 48-bit signed value returned in three consecutive 16-bit registers. R3*10,000^2 + R2*10,000 + R1.  Each registers range is -9,999 to +9,999
	TypeMod10_u64 DataType = "mod10-u64" // TODO: (not yet implemented) - MOD10k unsigned 64bit (2 registers) - A 64-bit unsigned value returned in four consecutive 16-bit registers. R4*10,000^3 + R3*10,000^2 + R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10_s64 DataType = "mod10-s64" // TODO: (not yet implemented) - MOD10k signed 64bit (2 registers) - A 64-bit signed value returned in four consecutive 16-bit registers. R4*10,000^3 + R3*10,000^2 + R2*10,000 + R1.  Each registers range is -9,999 to +9,999
)

type ByteOrder string

const (
	ByteOrderLebBew ByteOrder = "leb_bew" // LITTLE_ENDIAN, HIGH_WORD_FIRST
	ByteOrderLebLew ByteOrder = "leb_lew"
	ByteOrderBebLew ByteOrder = "beb_lew"
	ByteOrderBebBew ByteOrder = "beb_bew"
)

type IOType string

const (
	IOTypeRAW           IOType = "raw"
	IOTypeDigital       IOType = "digital"
	IOTypeAToDigital    IOType = "a_to_digital"
	IOTypeVoltageDC     IOType = "voltage_dc"
	IOTypeVoltage010    IOType = "voltage_0_10dc"
	IOTypeCurrent       IOType = "current"
	IOTypeThermistor    IOType = "thermistor"
	IOTypeThermistor10K IOType = "thermistor_10k_type_2"
)

type PointPriorityArrayMode string

const (
	PriorityArrayToPresentValue     PointPriorityArrayMode = "priority_array_to_present_value"      // This is a normal point type
	PriorityArrayToWriteValue       PointPriorityArrayMode = "priority_array_to_write_value"        // This is a point like a modbus writeable point which gets its present value from the read value, not directly from the priority array.
	ReadOnlyNoPriorityArrayRequired PointPriorityArrayMode = "read_only_no_priority_array_required" // This is a point like a modbus read only point, which doesn't need a priority array, only a present value
)

type ReadWriteType string

const (
	WriteToPriorityArray   ReadWriteType = "write_to_priority_array"
	ReadWritePriorityArray ReadWriteType = "read_write_to_priority_array"
	ReadPriorityArray      ReadWriteType = "read_priority_array"

	WriteToPresentValue   ReadWriteType = "write_to_present_value"
	ReadPresentValue      ReadWriteType = "read_present_value"
	ReadWritePresentValue ReadWriteType = "read_write_to_present_value"
)

type WriteMode string

const (
	ReadOnce          WriteMode = "read_once"            // Only Read Point Value Once.
	ReadOnly          WriteMode = "read_only"            // Only Read Point Value (poll rate defined by setting).
	WriteOnce         WriteMode = "write_once"           // Write the value on COV, don't Read.
	WriteOnceReadOnce WriteMode = "write_once_read_once" // Write the value on COV, Read Once.
	WriteAlways       WriteMode = "write_always"         // Write the value on every poll (poll rate defined by setting).
	WriteOnceThenRead WriteMode = "write_once_then_read" // Write the value on COV, then Read on each poll (poll rate defined by setting).
	WriteAndMaintain  WriteMode = "write_and_maintain"   // Write the value on COV, then Read on each poll (poll rate defined by setting). If the Read value does not match the Write value, Write the value again.
)

type PollPriority string

const (
	PRIORITY_ASAP   PollPriority = "asap"
	PRIORITY_HIGH   PollPriority = "high"
	PRIORITY_NORMAL PollPriority = "normal"
	PRIORITY_LOW    PollPriority = "low"
)

type PointState string

const (
	PointStatePollFailed       PointState = "poll-failed"
	PointStatePollOk           PointState = "poll-ok"
	PointStateWriteOk          PointState = "write-ok"
	PointStateApiWritePending  PointState = "api-write-pending"
	PointStateApiUpdatePending PointState = "api-update-pending"
)

type PollRate string

const (
	RATE_FAST   PollRate = "fast"
	RATE_NORMAL PollRate = "normal"
	RATE_SLOW   PollRate = "slow"
)

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
	PresentValue           *float64               `json:"present_value"` // point value, read only
	OriginalValue          *float64               `json:"original_value"`
	DisplayValue           *string                `json:"display_value"`        // 248.67 °F
	WriteValue             *float64               `json:"write_value"`          // writeValue was added so if user wanted to do a math function on the point write
	WriteValueOriginal     *float64               `json:"write_value_original"` // writeValue was added so if user wanted to do a math function on the point write
	CurrentPriority        *int                   `json:"current_priority,omitempty"`
	WritePriority          *int                   `json:"write_priority,omitempty"` // used for user just to select the write priority on example for bacnet
	IsOutput               *bool                  `json:"is_output"`                // used for as example: for bacnet-server we only support AV so if a point IsOutput = false then for mapping its set as a consumer but if true then its set as a prodcuer
	IsTypeBool             *bool                  `json:"is_type_bool"`
	InSync                 *bool                  `json:"in_sync"` // is set to false when a new value is written from the user example: if its false then modbus would write the new value. if user edits the point it will disable the COV for one time
	Fallback               *float64               `json:"fallback"`
	DeviceUUID             string                 `json:"device_uuid,omitempty" gorm:"type:string references devices;not null;default:null;uniqueIndex:idx_points_name_device_uuid"`
	EnableWriteable        *bool                  `json:"writeable,omitempty"`             // TODO: UI should hide the `write` action if not enabled
	MathOnPresentValue     string                 `json:"math_on_present_value,omitempty"` // TODO: THIS SHOULD BE DELETED WHEN SAFE TO DO SO.
	MathOnWriteValue       string                 `json:"math_on_write_value,omitempty"`   // TODO: THIS SHOULD BE DELETED WHEN SAFE TO DO SO.
	COV                    *float64               `json:"cov"`
	ObjectType             string                 `json:"object_type,omitempty"` // binaryInput, coil, if type os input don't return the priority array
	ObjectId               *int                   `json:"object_id,omitempty"`
	DataType               string                 `json:"data_type,omitempty"`       // int16, uint16, float32
	IsBitwise              *bool                  `json:"is_bitwise,omitempty"`      // true indicates data type is numeric, but we are taking only a specific bit of the number. Used for Modbus and Bacnet.
	BitwiseIndex           *int                   `json:"bitwise_index,omitempty"`   // if IsBitwise is set, then BitwiseIndex sets the index to read from numeric value
	ObjectEncoding         string                 `json:"object_encoding,omitempty"` // BEB_LEW bebLew
	IoNumber               string                 `json:"io_number,omitempty"`       // DI1,UI1,AO1, temp, pulse, motion
	IoType                 string                 `json:"io_type,omitempty"`         // 0-10dc, 0-40ma, thermistor
	AddressID              *int                   `json:"address_id"`                // for example a modbus address or bacnet address
	AddressLength          *int                   `json:"address_length"`            // for example a modbus address offset
	AddressUUID            *string                `json:"address_uuid,omitempty"`    // for example a droplet id (so a string)
	NextAvailableAddress   *bool                  `json:"use_next_available_address,omitempty"`
	Decimal                *uint32                `json:"decimal,omitempty"`
	MultiplicationFactor   *float64               `json:"multiplication_factor"`
	ScaleEnable            *bool                  `json:"scale_enable"`
	ScaleInMin             *float64               `json:"scale_in_min"`
	ScaleInMax             *float64               `json:"scale_in_max"`
	ScaleOutMin            *float64               `json:"scale_out_min"`
	ScaleOutMax            *float64               `json:"scale_out_max"`
	Offset                 *float64               `json:"offset"`
	UnitType               *string                `json:"unit_type,omitempty"` // temperature
	Unit                   *string                `json:"unit,omitempty"`
	UnitTo                 *string                `json:"unit_to,omitempty"` // with take the unit and convert to, this would affect the presentValue and the original value will be stored in the raw
	Priority               *Priority              `json:"priority,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Tags                   []*Tag                 `json:"tags,omitempty" gorm:"many2many:points_tags;constraint:OnDelete:CASCADE"`
	ValueUpdatedFlag       *bool                  `json:"value_updated_flag,omitempty"`      // This is used when a plugin updates the PresentValue (not from priority array) and it triggers UpdatePointValue() to broadcast to producers. Should only be set to FALSE from UpdatePointValue().
	PointPriorityArrayMode PointPriorityArrayMode `json:"point_priority_use_type,omitempty"` // This configures how the point handles the priority array and present value.
	WriteMode              WriteMode              `json:"write_mode,omitempty"`
	ReadWriteType          ReadWriteType          `json:"read_write_type,omitempty"`
	WritePollRequired      *bool                  `json:"write_required,omitempty"`
	ReadPollRequired       *bool                  `json:"read_required,omitempty"`
	PollPriority           PollPriority           `json:"poll_priority"`
	PollRate               PollRate               `json:"poll_rate"`
	PointState             PointState             `json:"point_state,omitempty"`
	BACnetWriteToPV        *bool                  `json:"bacnet_write_to_pv,omitempty"`
	HistoryConfig
	MetaTags          []*PointMetaTag `json:"meta_tags,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Connection        string          `json:"connection" gorm:"default:Connected"`
	ConnectionMessage *string         `json:"connection_message" gorm:""`
	PointHistories    []*PointHistory `json:"point_histories,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	CommonSourceUUID
	LastHistoryValue *float64 `json:"last_history_value"`
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
	PollState    PointState           `json:"poll_state"`
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
