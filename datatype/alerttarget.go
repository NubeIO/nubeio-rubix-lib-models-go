package datatype

import (
	"fmt"
	"strings"
)

type AlertTarget string

const (
	AlertTargetMobile AlertTarget = "mobile"
	AlertTargetNone   AlertTarget = "none"
)

var AlertTargetMap = map[AlertTarget]struct{}{
	AlertTargetMobile: {},
	AlertTargetNone:   {},
}

func (dt *AlertTarget) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := AlertTargetMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(AlertTargetMap))
	for m := range AlertTargetMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid alert target, i.e.: %s", strings.Join(v, " or "))
}
