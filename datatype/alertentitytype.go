package datatype

import (
	"fmt"
	"strings"
)

type AlertEntityType string

const (
	AlertEntityTypeGateway AlertEntityType = "gateway"
	AlertEntityTypeNetwork AlertEntityType = "network"
	AlertEntityTypeDevice  AlertEntityType = "device"
	AlertEntityTypePoint   AlertEntityType = "point"
	AlertEntityTypeService AlertEntityType = "service"
)

var AlertEntityTypeMap = map[AlertEntityType]struct{}{
	AlertEntityTypeGateway: {},
	AlertEntityTypeNetwork: {},
	AlertEntityTypeDevice:  {},
	AlertEntityTypePoint:   {},
	AlertEntityTypeService: {},
}

func (dt *AlertEntityType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := AlertEntityTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(AlertEntityTypeMap))
	for m := range AlertEntityTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid alert entity type, i.e.: %s", strings.Join(v, " or "))
}
