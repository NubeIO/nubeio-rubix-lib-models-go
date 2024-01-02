package datatype

import (
	"fmt"
	"strings"
)

type MemberState string

const (
	MemberStateVerified   MemberState = "VERIFIED"
	MemberStateUnVerified MemberState = "UNVERIFIED"
)

var MemberStateMap = map[MemberState]struct{}{
	MemberStateVerified:   {},
	MemberStateUnVerified: {},
}

func (dt *MemberState) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := MemberStateMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(MemberStateMap))
	for m := range MemberStateMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid member state, i.e.: %s", strings.Join(v, " or "))
}
