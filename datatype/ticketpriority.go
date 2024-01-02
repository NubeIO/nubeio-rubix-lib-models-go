package datatype

import (
	"fmt"
	"strings"
)

type TicketPriority string

const (
	TicketPriorityLow      TicketPriority = "LOW"
	TicketPriorityMedium   TicketPriority = "MEDIUM"
	TicketPriorityHigh     TicketPriority = "HIGH"
	TicketPriorityCritical TicketPriority = "CRITICAL"
)

var TicketPriorityMap = map[TicketPriority]struct{}{
	TicketPriorityLow:      {},
	TicketPriorityMedium:   {},
	TicketPriorityHigh:     {},
	TicketPriorityCritical: {},
}

func (dt *TicketPriority) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := TicketPriorityMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(TicketPriorityMap))
	for m := range TicketPriorityMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid ticket priority, i.e.: %s", strings.Join(v, " or "))
}
