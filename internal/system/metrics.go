package system

type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	NetworkIn   float64 `json:"network_in"`
	NetworkOut  float64 `json:"network_out"`
}

func GetSystemStats() (*SystemStats, error) {
	return &SystemStats{}, nil
}
