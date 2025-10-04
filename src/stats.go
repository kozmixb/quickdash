package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/disk"
)

const GIGABYTE float64 = 1024 * 1024 * 1024

func homeController(c echo.Context) error {
	mountPath := "C:\\" // linux = "/", windows = "C:\\"
	usage, err := disk.Usage(mountPath)
	if err != nil {
		log.Fatalf("Error getting disk usage for %s: %v", mountPath, err)
	}

	return c.Render(200, "index", map[string]interface{}{
		"cpu_count":           runtime.NumCPU(),
		"cpu_freq":            600.,
		"cpu_mem_avail":       22.4,
		"cpu_mem_free":        0.11,
		"cpu_mem_total":       232,
		"cpu_mem_used":        0.43,
		"cpu_percent":         1.8,
		"disk_usage_free":     fmt.Sprintf("%.2f", (float64(usage.Free) / GIGABYTE)),
		"disk_usage_percent":  fmt.Sprintf("%.2f", usage.UsedPercent),
		"disk_usage_total":    fmt.Sprintf("%.2f", (float64(usage.Total) / GIGABYTE)),
		"disk_usage_used":     fmt.Sprintf("%.2f", (float64(usage.Used) / GIGABYTE)),
		"sensor_temperatures": 52.6,
	})
}

func routes(e *echo.Echo) {
	e.GET("/", homeController)
}
