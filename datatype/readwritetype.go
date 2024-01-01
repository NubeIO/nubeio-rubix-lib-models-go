package datatype

import (
	"fmt"
	"strings"
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

var ReadWriteTypeMap = map[ReadWriteType]struct{}{
	WriteToPriorityArray:   {},
	ReadWritePriorityArray: {},
	ReadPriorityArray:      {},
	WriteToPresentValue:    {},
	ReadPresentValue:       {},
	ReadWritePresentValue:  {},
}

func (dt *ReadWriteType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ReadWriteTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ReadWriteTypeMap))
	for m := range ReadWriteTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid read write type, i.e.: %s", strings.Join(v, " or "))
}
