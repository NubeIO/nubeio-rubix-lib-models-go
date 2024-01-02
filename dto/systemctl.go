package dto

type SystemCtlProperty struct {
	Property string `json:"property"`
}

type SystemCtlStatus struct {
	Status string `json:"status"`
}

type SystemCtlState struct {
	State bool `json:"state"`
}

type SystemCtlStateStatus struct {
	State  bool   `json:"state"`
	Status string `json:"status"`
}
