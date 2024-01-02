package datatype

import (
	"fmt"
	"strings"
)

type AlertSeverity string

const (
	AlertSeverityCrucial AlertSeverity = "crucial"
	AlertSeverityMinor   AlertSeverity = "minor"
	AlertSeverityInfo    AlertSeverity = "info"
	AlertSeverityWarning AlertSeverity = "warning"
)

var AlertSeverityMap = map[AlertSeverity]struct{}{
	AlertSeverityCrucial: {},
	AlertSeverityMinor:   {},
	AlertSeverityInfo:    {},
	AlertSeverityWarning: {},
}

func (dt *AlertSeverity) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := AlertSeverityMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(AlertSeverityMap))
	for m := range AlertSeverityMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid alert severity, i.e.: %s", strings.Join(v, " or "))
}
