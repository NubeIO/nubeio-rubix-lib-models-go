package dto

type MemoryUsage struct {
	MemoryPercentageUsed float64 `json:"memory_percentage_used"`
	MemoryPercentage     string  `json:"memory_percentage"`
	MemoryAvailable      string  `json:"memory_available"`
	MemoryFree           string  `json:"memory_free"`
	MemoryUsed           string  `json:"memory_used"`
	MemoryTotal          string  `json:"memory_total"`
	SwapPercentageUsed   float64 `json:"swap_percentage_used"`
	SwapPercentage       string  `json:"swap_percentage"`
	SwapFree             string  `json:"swap_free"`
	SwapUsed             string  `json:"swap_used"`
	SwapTotal            string  `json:"swap_total"`
}
