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
	mem, _ := stats.ReadMemInfo()
	host, _ := stats.ReadHostInfo()

	return c.Render(200, "index", map[string]interface{}{
		"host_ip":             host.IP,
		"host_name":           host.Hostname,
		"host_arch":           host.Arch,
		"host_uptime":         host.Uptime,
		"host_os":             host.OS,
		"host_platform":       host.Platform,
		"cpu_count":           cpu.Cores,
		"cpu_freq":            cpu.Mhz,
		"cpu_percent":         cpu.Percent,
		"sensor_temperatures": cpu.Temp,
		"cpu_mem_avail":       mem.Available,
		"cpu_mem_free":        mem.Free,
		"cpu_mem_total":       mem.Total,
		"cpu_mem_used":        mem.Percent,
		"disk_usage_free":     disk.Free,
		"disk_usage_percent":  disk.Percent,
		"disk_usage_total":    disk.Total,
		"disk_usage_used":     disk.Used,
	})
}

func routes(e *echo.Echo) {
	e.GET("/", homeController)
}
