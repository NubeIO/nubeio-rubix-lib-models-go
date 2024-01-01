package datatype

import (
	"fmt"
	"strings"
)

type HistoryType string

const (
	HistoryTypeCov            HistoryType = "COV"
	HistoryTypeInterval       HistoryType = "INTERVAL"
	HistoryTypeCovAndInterval HistoryType = "COV_AND_INTERVAL"
)

var HistoryTypeMap = map[HistoryType]struct{}{
	HistoryTypeCov:            {},
	HistoryTypeInterval:       {},
	HistoryTypeCovAndInterval: {},
}

var HistoryTypeCovMap = map[HistoryType]struct{}{
	HistoryTypeCov:            {},
	HistoryTypeCovAndInterval: {},
}

func (dt *HistoryType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := HistoryTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(HistoryTypeMap))
	for m := range HistoryTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid history type, i.e.: %s", strings.Join(v, " or "))
}

func (dt *HistoryType) IsCov() bool {
	if dt == nil {
		return false
	}
	if _, ok := HistoryTypeCovMap[*dt]; ok {
		return true
	}
	return false
}
