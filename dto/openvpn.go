package dto

type OpenVPNClient struct {
	VirtualIP      string `json:"virtual_ip"`
	ReceivedBytes  int    `json:"received_bytes"`
	SentBytes      int    `json:"sent_bytes"`
	ConnectedSince string `json:"connected_since"`
}

type OpenVPNBody struct {
	Name string `json:"name"`
}

type OpenVPNConfig struct {
	Data string `json:"data"`
}

type OpenVPNCredentials struct {
	IPAddress string `json:"ip_address"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Server    string `json:"server"`
}
