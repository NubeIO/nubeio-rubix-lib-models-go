package datatype

import (
	"fmt"
	"strings"
)

type ByteOrder string

const (
	ByteOrderLebBew ByteOrder = "leb_bew" // LITTLE_ENDIAN, HIGH_WORD_FIRST
	ByteOrderLebLew ByteOrder = "leb_lew"
	ByteOrderBebLew ByteOrder = "beb_lew"
	ByteOrderBebBew ByteOrder = "beb_bew"
)

var ByteOrderMap = map[ByteOrder]struct{}{
	ByteOrderLebBew: {},
	ByteOrderLebLew: {},
	ByteOrderBebLew: {},
	ByteOrderBebBew: {},
}

func (dt *ByteOrder) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ByteOrderMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ByteOrderMap))
	for m := range ByteOrderMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid byte order, i.e.: %s", strings.Join(v, " or "))
}
