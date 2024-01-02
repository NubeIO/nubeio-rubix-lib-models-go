package datatype

import (
	"fmt"
	"strings"
)

type EmailConfigType string

const (
	EmailConfigTypeSmtp EmailConfigType = "SMTP"
	EmailConfigTypeSes  EmailConfigType = "SES"
)

var EmailConfigTypeMap = map[EmailConfigType]struct{}{
	EmailConfigTypeSmtp: {},
	EmailConfigTypeSes:  {},
}

func (dt *EmailConfigType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := EmailConfigTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(EmailConfigTypeMap))
	for m := range EmailConfigTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid email config type, i.e.: %s", strings.Join(v, " or "))
}
