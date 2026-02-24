package system

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	NetworkIn   float64 `json:"network_in"`
	NetworkOut  float64 `json:"network_out"`
}

func GetSystemStats() (*SystemStats, error) {
	memoryInfo, err := mem.VirtualMemory()

	if err != nil {
		return nil, err
	}
	stats := &SystemStats{
		MemoryUsage: memoryInfo.UsedPercent,
	}
	return stats, nil
}
