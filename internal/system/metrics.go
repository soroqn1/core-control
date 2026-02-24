package system

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	NetworkIn   float64 `json:"network_in"`
	NetworkOut  float64 `json:"network_out"`
}

func GetSystemStats() (*SystemStats, error) {
	CPUUsage, err := cpu.Percent(500*time.Millisecond, false)
	memoryInfo, err := mem.VirtualMemory()
	//TODO: Network Usage

	if err != nil {
		return nil, err
	}
	stats := &SystemStats{
		MemoryUsage: memoryInfo.UsedPercent,
		CPUUsage:    CPUUsage[0],
	}
	return stats, nil
}
