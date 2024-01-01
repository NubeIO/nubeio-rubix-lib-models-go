package datatype

import (
	"fmt"
	"strings"
)

type PointState string

const (
	PointStatePollFailed       PointState = "poll-failed"
	PointStatePollOk           PointState = "poll-ok"
	PointStateWriteOk          PointState = "write-ok"
	PointStateApiWritePending  PointState = "api-write-pending"
	PointStateApiUpdatePending PointState = "api-update-pending"
)

var PointStateMap = map[PointState]struct{}{
	PointStatePollFailed:       {},
	PointStatePollOk:           {},
	PointStateWriteOk:          {},
	PointStateApiWritePending:  {},
	PointStateApiUpdatePending: {},
}

func (dt *PointState) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := PointStateMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(PointStateMap))
	for m := range PointStateMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid point state, i.e.: %s", strings.Join(v, " or "))
}
