package datatype

import (
	"fmt"
	"strings"
)

type EmailStatus string

const (
	EmailStatusSending EmailStatus = "Sending"
	EmailStatusSent    EmailStatus = "Sent"
	EmailStatusFailed  EmailStatus = "Failed"
)

var EmailStatusMap = map[EmailStatus]struct{}{
	EmailStatusSending: {},
	EmailStatusSent:    {},
	EmailStatusFailed:  {},
}

func (dt *EmailStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := EmailStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(EmailStatusMap))
	for m := range EmailStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid email status, i.e.: %s", strings.Join(v, " or "))
}
