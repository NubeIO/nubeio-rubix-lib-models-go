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
	TimeZone    string   `json:"time_zone"`
	Groups      []*Group `json:"groups" gorm:"constraint:OnDelete:CASCADE"`
	Views       []*View  `json:"views" gorm:"constraint:OnDelete:CASCADE"`
}
