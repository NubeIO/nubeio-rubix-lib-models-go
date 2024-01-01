package datatype

import (
	"fmt"
	"strings"
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

var WriteModeMap = map[WriteMode]struct{}{
	ReadOnce:          {},
	ReadOnly:          {},
	WriteOnce:         {},
	WriteOnceReadOnce: {},
	WriteAlways:       {},
	WriteOnceThenRead: {},
	WriteAndMaintain:  {},
}

func (dt *WriteMode) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := WriteModeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(WriteModeMap))
	for m := range WriteModeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid write mode, i.e.: %s", strings.Join(v, " or "))
}
