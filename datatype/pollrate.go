package datatype

import (
	"fmt"
	"strings"
)

type PollRate string

const (
	RateFast   PollRate = "fast"
	RateNormal PollRate = "normal"
	RateSlow   PollRate = "slow"
)

var PollRateMap = map[PollRate]struct{}{
	RateFast:   {},
	RateNormal: {},
	RateSlow:   {},
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
