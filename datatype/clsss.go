package datatype

import (
	"fmt"
	"strings"
)

type Class string

const (
	ClassPoint    Class = "point"
	ClassSchedule Class = "schedule"
	ClassQuery    Class = "query"
)

var ClassMap = map[Class]int8{
	ClassPoint:    0,
	ClassSchedule: 0,
	ClassQuery:    0,
}

func (dt *Class) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ClassMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ClassMap))
	for m := range ClassMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid class, i.e.: %s", strings.Join(v, " or "))
}
