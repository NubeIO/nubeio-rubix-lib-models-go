package datatype

import (
	"fmt"
	"strings"
)

type AlertType string

const (
	AlertTypePing      AlertType = "ping"
	AlertTypeFault     AlertType = "fault"
	AlertTypeThreshold AlertType = "threshold"
	AlertTypeFlatLine  AlertType = "flat-line"
)

var AlertTypeMap = map[AlertType]struct{}{
	AlertTypePing:      {},
	AlertTypeFault:     {},
	AlertTypeThreshold: {},
	AlertTypeFlatLine:  {},
}

func (dt *AlertType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := AlertTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(AlertTypeMap))
	for m := range AlertTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid alert type, i.e.: %s", strings.Join(v, " or "))
}
