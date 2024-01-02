package datatype

import (
	"fmt"
	"strings"
)

type StreamLogStatus string

const (
	StreamLogCreating StreamLogStatus = "Creating"
	StreamLogCreated  StreamLogStatus = "Created"
	StreamLogFailed   StreamLogStatus = "Failed"
)

var StreamLogStatusMap = map[StreamLogStatus]struct{}{
	StreamLogCreating: {},
	StreamLogCreated:  {},
	StreamLogFailed:   {},
}

func (dt *StreamLogStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := StreamLogStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(StreamLogStatusMap))
	for m := range StreamLogStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid stream log status, i.e.: %s", strings.Join(v, " or "))
}
