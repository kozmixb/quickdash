package main

import (
	"quickdash/src/stats"

	"fmt"

	"github.com/labstack/echo/v4"
)

func homeController(c echo.Context) error {

	cpu, err := stats.ReadCPUInfo()
	if err != nil {
		fmt.Println(err)
	}

	disk, _ := stats.ReadDiskInfo()

	return c.Render(200, "index", map[string]interface{}{
		"cpu_count":           cpu.Cores,
		"cpu_freq":            cpu.Mhz,
		"cpu_percent":         cpu.Percent,
		"sensor_temperatures": cpu.Temp,
		"cpu_mem_avail":       22.4,
		"cpu_mem_free":        0.11,
		"cpu_mem_total":       232,
		"cpu_mem_used":        0.43,
		"disk_usage_free":     disk.Free,
		"disk_usage_percent":  disk.Percent,
		"disk_usage_total":    disk.Total,
		"disk_usage_used":     disk.Used,
	})
}

func routes(e *echo.Echo) {
	e.GET("/", homeController)
}
