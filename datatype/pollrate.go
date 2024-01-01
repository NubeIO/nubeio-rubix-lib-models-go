package datatype

import (
	"fmt"
	"strings"
)

type PollRate string

const (
	RATE_FAST   PollRate = "fast"
	RATE_NORMAL PollRate = "normal"
	RATE_SLOW   PollRate = "slow"
)

var PollRateMap = map[PollRate]struct{}{
	RATE_FAST:   {},
	RATE_NORMAL: {},
	RATE_SLOW:   {},
}

func (dt *PollRate) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := PollRateMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(PollRateMap))
	for m := range PollRateMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid poll rate, i.e.: %s", strings.Join(v, " or "))
}
