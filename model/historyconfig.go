package model

// HistoryConfig for configuring point histories
type HistoryConfig struct {
	CommonHistoryEnable
	HistoryType         HistoryType `json:"history_type,omitempty"`
	HistoryInterval     *int        `json:"history_interval,omitempty"`
	HistoryCOVThreshold *float64    `json:"history_cov_threshold,omitempty"`
}

// HistoryType types of histories
type HistoryType string

const (
	HistoryTypeCov            HistoryType = "COV"
	HistoryTypeInterval       HistoryType = "INTERVAL"
	HistoryTypeCovAndInterval HistoryType = "COV_AND_INTERVAL"
)

var HistoryTypeMap = map[HistoryType]int8{
	HistoryTypeCov:            0,
	HistoryTypeInterval:       0,
	HistoryTypeCovAndInterval: 0,
}

var HistoryTypeCovMap = map[HistoryType]int8{
	HistoryTypeCov:            0,
	HistoryTypeCovAndInterval: 0,
}