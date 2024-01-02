package datatype

import (
	"fmt"
	"strings"
)

type DataType string

const (
	TypeDigital  DataType = "digital"
	TypeInt16    DataType = "int16"
	TypeUint16   DataType = "uint16"
	TypeInt32    DataType = "int32"
	TypeUint32   DataType = "uint32"
	TypeInt64    DataType = "int64"
	TypeUint64   DataType = "uint64"
	TypeFloat32  DataType = "float32"
	TypeFloat64  DataType = "float64"
	TypeMod10U32 DataType = "mod10-u32" // MOD10k unsigned 32bit (2 registers) - A 32-bit unsigned value returned in two consecutive 16-bit registers. R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10S32 DataType = "mod10-s32" // TODO: (not yet implemented) - MOD10k signed 32bit (2 registers) - A 32-bit signed value returned in two consecutive 16-bit registers. R2*10,000 + R1.  Each registers range is -9,999 to +9,999
	TypeMod10U48 DataType = "mod10-u48" // TODO: (not yet implemented) - MOD10k unsigned 48bit (2 registers) - A 48-bit unsigned value returned in three consecutive 16-bit registers. R3*10,000^2 + R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10S48 DataType = "mod10-s48" // TODO: (not yet implemented) - MOD10k signed 48bit (2 registers) - A 48-bit signed value returned in three consecutive 16-bit registers. R3*10,000^2 + R2*10,000 + R1.  Each registers range is -9,999 to +9,999
	TypeMod10U64 DataType = "mod10-u64" // TODO: (not yet implemented) - MOD10k unsigned 64bit (2 registers) - A 64-bit unsigned value returned in four consecutive 16-bit registers. R4*10,000^3 + R3*10,000^2 + R2*10,000 + R1.  Each registers range is 0 to +9,999
	TypeMod10S64 DataType = "mod10-s64" // TODO: (not yet implemented) - MOD10k signed 64bit (2 registers) - A 64-bit signed value returned in four consecutive 16-bit registers. R4*10,000^3 + R3*10,000^2 + R2*10,000 + R1.  Each registers range is -9,999 to +9,999
)

var DataTypeMap = map[DataType]struct{}{
	TypeDigital:  {},
	TypeInt16:    {},
	TypeUint16:   {},
	TypeInt32:    {},
	TypeUint32:   {},
	TypeInt64:    {},
	TypeUint64:   {},
	TypeFloat32:  {},
	TypeFloat64:  {},
	TypeMod10U32: {},
	TypeMod10S32: {},
	TypeMod10U48: {},
	TypeMod10S48: {},
	TypeMod10U64: {},
	TypeMod10S64: {},
}

func (dt *DataType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := DataTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(DataTypeMap))
	for m := range DataTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid data type, i.e.: %s", strings.Join(v, " or "))
}
