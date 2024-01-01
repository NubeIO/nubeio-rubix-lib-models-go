package datatype

import (
	"fmt"
	"strings"
)

type PollPriority string

const (
	PRIORITY_ASAP   PollPriority = "asap"
	PRIORITY_HIGH   PollPriority = "high"
	PRIORITY_NORMAL PollPriority = "normal"
	PRIORITY_LOW    PollPriority = "low"
)

var PollPriorityMap = map[PollPriority]struct{}{
	PRIORITY_ASAP:   {},
	PRIORITY_HIGH:   {},
	PRIORITY_NORMAL: {},
	PRIORITY_LOW:    {},
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
