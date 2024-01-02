package datatype

import (
	"fmt"
	"strings"
)

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

var ObjectTypesMap = map[ObjectType]struct{}{
	// bacnet
	ObjTypeAnalogInput:  {},
	ObjTypeAnalogOutput: {},
	ObjTypeAnalogValue:  {},
	ObjTypeBinaryInput:  {},
	ObjTypeBinaryOutput: {},
	ObjTypeBinaryValue:  {},
	ObjAnalogInput:      {},
	ObjAnalogOutput:     {},
	ObjAnalogValue:      {},
	ObjBinaryInput:      {},
	ObjBinaryOutput:     {},
	ObjBinaryValue:      {},
	ObjTypeEnumInput:    {},
	ObjTypeEnumOutput:   {},
	ObjTypeEnumValue:    {},
	// modbus
	ObjTypeReadCoil:           {},
	ObjTypeReadCoils:          {},
	ObjTypeReadDiscreteInput:  {},
	ObjTypeReadDiscreteInputs: {},
	ObjTypeWriteCoil:          {},
	ObjTypeWriteCoils:         {},
	ObjTypeReadRegister:       {},
	ObjTypeReadRegisters:      {},
	ObjTypeReadHolding:        {},
	ObjTypeReadHoldings:       {},
	ObjTypeWriteHolding:       {},
	ObjTypeWriteHoldings:      {},
	ObjTypeReadInt16:          {},
	ObjTypeWriteInt16:         {},
	ObjTypeReadUint16:         {},
	ObjTypeWriteUint16:        {},
	ObjTypeReadInt32:          {},
	ObjTypeReadUint32:         {},
	ObjTypeReadFloat32:        {},
	ObjTypeWriteFloat32:       {},
	ObjTypeReadFloat64:        {},
	ObjTypeWriteFloat64:       {},
	ObjTypeCoil:               {},
	ObjTypeDiscreteInput:      {},
	ObjTypeInputRegister:      {},
	ObjTypeHoldingRegister:    {},
}

func (dt *ObjectType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ObjectTypesMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ObjectTypesMap))
	for m := range ObjectTypesMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a object type, i.e.: %s", strings.Join(v, " or "))
}
