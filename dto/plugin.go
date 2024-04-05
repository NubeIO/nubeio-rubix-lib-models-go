package dto

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
)

type Plugin struct {
	Name      string `json:"name"`
	Arch      string `json:"arch"`
	Version   string `json:"version,omitempty"`
	Extension string `json:"extension"`
}

// PluginExternal holds information about a plugin instance for one user.
type PluginExternal struct {
	UUID         string               `json:"uuid" sql:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Name         string               `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Enabled      bool                 `json:"enabled"`
	Network      model.Network        `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
	Author       string               `json:"author,omitempty" form:"author" query:"author"`
	Website      string               `json:"website,omitempty" form:"website" query:"website"`
	License      string               `json:"license,omitempty" form:"license" query:"license"`
	HasNetwork   bool                 `json:"has_network"`
	Capabilities []string             `json:"capabilities,omitempty"`
	MessageLevel string               `json:"message_level"`
	Message      string               `json:"message"`
	State        datatype.PluginState `json:"state"`
}
