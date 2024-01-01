package model

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

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
	Topic   string       `json:"string"`
	Qos     datatype.QOS `json:"qos"`
	Retain  bool         `json:"retain"`
	Payload string       `json:"payload"`
}
