package datatype

import (
	"fmt"
	"strings"
)

type PluginState string

const (
	PluginStateRunning PluginState = "active"
	PluginStateFailed  PluginState = "failed"
)

var PluginStateMap = map[PluginState]struct{}{
	PluginStateRunning: {},
	PluginStateFailed:  {},
}

func (dt *PluginState) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := PluginStateMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(PluginStateMap))
	for m := range PluginStateMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid plugin state, i.e.: %s", strings.Join(v, " or "))
}
