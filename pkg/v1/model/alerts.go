package model

import "time"

type AlertCategory struct {
	Type string // loss of data, offline
}

type AlertType struct {
	Type string // point, device
}

var CommonAlertTypes = struct {
	HostPing      string
	DeviceOffline string
}{
	HostPing:      "system_ping",
	DeviceOffline: "device_offline",
}

type Alert struct {
	UUID      string    `json:"uuid" gorm:"primaryKey"`
	From      string    `json:"from"`
	HostUUID  string    `json:"host_uuid"`
	Host      string    `json:"host"`
	AlertType string    `json:"alert_type"`
	Count     uint      `json:"count"`
	Date      time.Time `json:"date"`
}

type Message struct {
	ID      uint      `json:"uuid" gorm:"primaryKey"`
	Title   string    `json:"title,omitempty"`
	Message string    `json:"message,omitempty"`
	Type    string    `json:"type,omitempty"`
	Date    time.Time `json:"date,omitempty"`
}
