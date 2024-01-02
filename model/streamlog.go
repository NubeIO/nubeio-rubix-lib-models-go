package model

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"gorm.io/datatypes"
)

type StreamLog struct {
	CommonUUID
	Service        string                   `json:"service,omitempty"`
	Duration       int                      `json:"duration,omitempty"`
	NumberOfLines  int                      `json:"number_of_lines,omitempty"`
	Status         datatype.StreamLogStatus `json:"status,omitempty"`
	KeyWordsFilter datatypes.JSON           `json:"key_words_filter,omitempty"`
	Messages       datatypes.JSON           `json:"messages,omitempty"`
}
