package datatype

import (
	"fmt"
	"strings"
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

var IOTypeMap = map[IOType]struct{}{
	IOTypeRAW:           {},
	IOTypeDigital:       {},
	IOTypeAToDigital:    {},
	IOTypeVoltageDC:     {},
	IOTypeVoltage010:    {},
	IOTypeCurrent:       {},
	IOTypeThermistor:    {},
	IOTypeThermistor10K: {},
}

func (dt *IOType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := IOTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(IOTypeMap))
	for m := range IOTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid IO type, i.e.: %s", strings.Join(v, " or "))
}
