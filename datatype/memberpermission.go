package datatype

import (
	"fmt"
	"strings"
)

type MemberPermission string

const (
	MemberPermissionRead  MemberPermission = "READ"
	MemberPermissionWrite MemberPermission = "WRITE"
)

var MemberPermissionMap = map[MemberPermission]struct{}{
	MemberPermissionRead:  {},
	MemberPermissionWrite: {},
}

func (dt *MemberPermission) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if *dt == "" {
		*dt = MemberPermissionRead
		return nil
	}
	if _, ok := MemberPermissionMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(MemberPermissionMap))
	for m := range MemberPermissionMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid member permission, i.e.: %s", strings.Join(v, " or "))
}
