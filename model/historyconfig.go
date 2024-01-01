package model

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

// HistoryConfig for configuring point histories
type HistoryConfig struct {
	CommonHistoryEnable
	HistoryType         datatype.HistoryType `json:"history_type,omitempty"`
	HistoryInterval     *int                 `json:"history_interval,omitempty"`
	HistoryCOVThreshold *float64             `json:"history_cov_threshold,omitempty"`
}
