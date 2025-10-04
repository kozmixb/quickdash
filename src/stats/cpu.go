package stats

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

type CPUStats struct {
	Cores   int
	Mhz     int
	Percent float64
	Temp    float64
}

func readCPUTemp() (float64, error) {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		return 0, fmt.Errorf("failed to get sensor temperatures: %w", err)
	}

	for _, t := range temps {
		if t.SensorKey == "cpu_thermal" || t.SensorKey == "soc_thermal" || t.SensorKey == "thermal_zone0" {
			return t.Temperature, nil
		}
	}

	return 0, fmt.Errorf("could not find a known CPU temperature sensor")
}

func readCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, fmt.Errorf("error getting CPU percent: %v", err)
	}

	return convertTo2Decimals(percentages[0]), nil
}

func ReadCPUInfo() (CPUStats, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return CPUStats{
			Cores: runtime.NumCPU(),
		}, err
	}

	percentage, err := readCPUUsage()
	if err != nil {
		return CPUStats{
			Cores: runtime.NumCPU(),
			Mhz:   int(cpuInfo[0].Mhz),
		}, err
	}

	temp, err := readCPUTemp()
	if err != nil {
		return CPUStats{
			Cores:   runtime.NumCPU(),
			Mhz:     int(cpuInfo[0].Mhz),
			Percent: percentage,
		}, nil
	}

	return CPUStats{
		Cores:   runtime.NumCPU(),
		Mhz:     int(cpuInfo[0].Mhz),
		Percent: percentage,
		Temp:    temp,
	}, nil
}
