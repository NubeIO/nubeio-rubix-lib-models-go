package model

import "gorm.io/datatypes"

type StreamLogStatus string

const (
	StreamLogCreating StreamLogStatus = "Creating"
	StreamLogCreated  StreamLogStatus = "Created"
	StreamLogFailed   StreamLogStatus = "Failed"
)

type StreamLog struct {
	CommonUUID
	Service        string          `json:"service,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	NumberOfLines  int             `json:"number_of_lines,omitempty"`
	Status         StreamLogStatus `json:"status,omitempty"`
	KeyWordsFilter datatypes.JSON  `json:"key_words_filter,omitempty"`
	Messages       datatypes.JSON  `json:"messages,omitempty"`
}
