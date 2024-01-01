package datatype

import (
	"fmt"
	"strings"
)

type CreateStatus string

const (
	Creating     CreateStatus = "Creating"
	Created      CreateStatus = "Created"
	CreateFailed CreateStatus = "Failed"
)

var CreateStatusMap = map[CreateStatus]struct{}{
	Creating:     {},
	Created:      {},
	CreateFailed: {},
}

func (dt *CreateStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := CreateStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(CreateStatusMap))
	for m := range CreateStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid craete status, i.e.: %s", strings.Join(v, " or "))
}
