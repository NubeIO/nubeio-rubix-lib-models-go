package datatype

import (
	"fmt"
	"strings"
)

type RestoreStatus string

const (
	Restoring     RestoreStatus = "Restoring"
	Restored      RestoreStatus = "Restored"
	RestoreFailed RestoreStatus = "Failed"
)

var RestoreStatusMap = map[RestoreStatus]struct{}{
	Restoring:     {},
	Restored:      {},
	RestoreFailed: {},
}

func (dt *RestoreStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := RestoreStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(RestoreStatusMap))
	for m := range RestoreStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid restore status, i.e.: %s", strings.Join(v, " or "))
}
