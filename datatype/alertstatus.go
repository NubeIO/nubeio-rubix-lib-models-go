package datatype

import (
	"fmt"
	"strings"
)

type AlertStatus string

const (
	AlertStatusActive       AlertStatus = "active"
	AlertStatusAcknowledged AlertStatus = "acknowledged"
	AlertStatusClosed       AlertStatus = "closed"
)

var AlertStatusMap = map[AlertStatus]struct{}{
	AlertStatusActive:       {},
	AlertStatusAcknowledged: {},
	AlertStatusClosed:       {},
}

func (dt *AlertStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := AlertStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(AlertStatusMap))
	for m := range AlertStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid alert status, i.e.: %s", strings.Join(v, " or "))
}
