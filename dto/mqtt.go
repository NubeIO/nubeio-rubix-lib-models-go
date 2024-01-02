package dto

import "github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"

type MqttBody struct {
	Topic   string       `json:"string"`
	Qos     datatype.QOS `json:"qos"`
	Retain  bool         `json:"retain"`
	Payload string       `json:"payload"`
}

type MqttPoint struct {
	NetworkName string       `json:"network_name,omitempty"`
	DeviceName  string       `json:"device_name,omitempty"`
	PointName   string       `json:"point_name,omitempty"`
	PointUUID   string       `json:"point_uuid,omitempty"`
	Priority    *PointWriter `json:"priority,omitempty"`
}
