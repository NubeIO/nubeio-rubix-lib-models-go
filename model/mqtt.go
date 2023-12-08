package model

type QOS byte

const (
	// AtMostOnce means the broker will deliver at most once
	AtMostOnce QOS = iota
	// AtLeastOnce means the broker will deliver c message at least once
	AtLeastOnce
	// ExactlyOnce means the broker will deliver c message exactly once
	ExactlyOnce
)

type MqttConnection struct {
	CommonUUID
	Enabled                       bool   `json:"enabled,omitempty"`
	Master                        bool   `json:"master,omitempty"`
	Name                          string `json:"name,omitempty"`
	Host                          string `json:"host,omitempty"`
	Port                          int    `json:"port,omitempty"`
	Authentication                bool   `json:"authentication,omitempty"`
	Username                      string `json:"username,omitempty"`
	Password                      string `json:"password,omitempty"`
	Keepalive                     int    `json:"keepalive,omitempty"`
	Qos                           int    `json:"qos,omitempty"`
	Retain                        bool   `json:"retain,omitempty"`
	AttemptReconnectOnUnavailable bool   `json:"attempt_reconnect_on_unavailable,omitempty"`
	AttemptReconnectSecs          int    `json:"attempt_reconnect_secs,omitempty"`
	Timeout                       int    `json:"timeout,omitempty"`
	IntegrationUUID               string `json:"integration_uuid" gorm:"type:varchar(255) references integrations;null;default:null"`
}

type MqttBody struct {
	Topic   string `json:"string"`
	Qos     QOS    `json:"qos"`
	Retain  bool   `json:"retain"`
	Payload string `json:"payload"`
}
