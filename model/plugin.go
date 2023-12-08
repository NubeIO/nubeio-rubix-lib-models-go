package model

// Plugin holds information about the plugin.
type Plugin struct {
	UUID    string `json:"uuid" sql:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Name    string `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Enabled bool   `json:"enabled"`
	Config  []byte
	Storage []byte
	Network Network `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
}

// PluginExternal holds information about a plugin instance for one user.
type PluginExternal struct {
	UUID         string   `json:"uuid" sql:"uuid"  gorm:"type:varchar(255);unique;primaryKey"`
	Name         string   `json:"name"  gorm:"type:varchar(255);unique;not null"`
	Enabled      bool     `json:"enabled"`
	Network      Network  `json:"networks" gorm:"constraint:OnDelete:CASCADE"`
	Author       string   `json:"author,omitempty" form:"author" query:"author"`
	Website      string   `json:"website,omitempty" form:"website" query:"website"`
	License      string   `json:"license,omitempty" form:"license" query:"license"`
	HasNetwork   bool     `json:"has_network"`
	Capabilities []string `json:"capabilities,omitempty"`
}
