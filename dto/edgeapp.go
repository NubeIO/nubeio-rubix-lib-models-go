package dto

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
)

type AppsStatus struct {
	Name        string                 `json:"name,omitempty"`
	Version     string                 `json:"version,omitempty"`
	ServiceName string                 `json:"service_name,omitempty"`
	State       *systemctl.SystemState `json:"state,omitempty"`
}
