package model

type Location struct {
	CommonUUID
	CommonNameUnique
	Description string   `json:"description"`
	Address     string   `json:"address"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	Zip         string   `json:"zip"`
	Country     string   `json:"country"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Timezone    string   `json:"timezone"`
	Groups      []*Group `json:"groups,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Views       []*View  `json:"views,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}
