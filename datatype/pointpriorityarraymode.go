package datatype

import (
	"fmt"
	"strings"
)

type PointPriorityArrayMode string

const (
	PriorityArrayToPresentValue     PointPriorityArrayMode = "priority_array_to_present_value"      // This is a normal point type
	PriorityArrayToWriteValue       PointPriorityArrayMode = "priority_array_to_write_value"        // This is a point like a modbus writeable point which gets its present value from the read value, not directly from the priority array.
	ReadOnlyNoPriorityArrayRequired PointPriorityArrayMode = "read_only_no_priority_array_required" // This is a point like a modbus read only point, which doesn't need a priority array, only a present value
)

var PointPriorityArrayModeMap = map[PointPriorityArrayMode]struct{}{
	PriorityArrayToPresentValue:     {},
	PriorityArrayToWriteValue:       {},
	ReadOnlyNoPriorityArrayRequired: {},
}

func (dt *PointPriorityArrayMode) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := PointPriorityArrayModeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(PointPriorityArrayModeMap))
	for m := range PointPriorityArrayModeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid point priority array mode, i.e.: %s", strings.Join(v, " or "))
}
