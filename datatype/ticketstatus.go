package datatype

import (
	"fmt"
	"strings"
)

type TicketStatus string

const (
	TicketStatusNew      TicketStatus = "NEW"
	TicketStatusReplied  TicketStatus = "REPLIED"
	TicketStatusResolved TicketStatus = "RESOLVED"
	TicketStatusClosed   TicketStatus = "CLOSED"
	TicketStatusBlocked  TicketStatus = "BLOCKED"
)

var TicketStatusMap = map[TicketStatus]struct{}{
	TicketStatusNew:      {},
	TicketStatusReplied:  {},
	TicketStatusResolved: {},
	TicketStatusClosed:   {},
	TicketStatusBlocked:  {},
}

func (dt *TicketStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := TicketStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(TicketStatusMap))
	for m := range TicketStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid ticket status, i.e.: %s", strings.Join(v, " or "))
}
