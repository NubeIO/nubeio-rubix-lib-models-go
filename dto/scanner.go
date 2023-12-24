package dto

type Scanner struct {
	Count int      `json:"count"`
	Iface string   `json:"iface"`
	Ip    string   `json:"ip"`
	Ports []string `json:"ports"`
}
