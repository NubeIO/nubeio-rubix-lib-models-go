package model

// HistoryConfig for configuring point histories
type HistoryConfig struct {
	HistoryEnable       *bool       `json:"history_enable,omitempty"`
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
