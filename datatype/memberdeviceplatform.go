package datatype

import (
	"fmt"
	"strings"
)

type MemberDevicePlatform string

const (
	MemberDevicePlatformAndroid MemberDevicePlatform = "ANDROID"
	MemberDevicePlatformIos     MemberDevicePlatform = "IOS"
)

var MemberDevicePlatformMap = map[MemberDevicePlatform]struct{}{
	MemberDevicePlatformAndroid: {},
	MemberDevicePlatformIos:     {},
}

func (dt *MemberDevicePlatform) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := MemberDevicePlatformMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(MemberDevicePlatformMap))
	for m := range MemberDevicePlatformMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid member device platform, i.e.: %s", strings.Join(v, " or "))
}
