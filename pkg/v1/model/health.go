package model

type Health struct {
	Health   string `json:"health"`
	Database string `json:"database"`
	ROS      bool   `json:"ros"`
}

const (
	// StatusGreen everything is alright.
	StatusGreen = "green"
	// StatusOrange some things are alright.
	StatusOrange = "orange"
	// StatusRed nothing is alright.
	StatusRed = "red"
)
