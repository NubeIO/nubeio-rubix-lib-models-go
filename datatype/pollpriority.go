package datatype

import (
	"fmt"
	"strings"
)

type PollPriority string

const (
	PriorityASAP   PollPriority = "asap"
	PriorityHigh   PollPriority = "high"
	PriorityNormal PollPriority = "normal"
	PriorityLow    PollPriority = "low"
)

var PollPriorityMap = map[PollPriority]struct{}{
	PriorityASAP:   {},
	PriorityHigh:   {},
	PriorityNormal: {},
	PriorityLow:    {},
}

func (dt *PollPriority) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := PollPriorityMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(PollPriorityMap))
	for m := range PollPriorityMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid poll priority, i.e.: %s", strings.Join(v, " or "))
}
