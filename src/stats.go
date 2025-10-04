package main

import (
	"runtime"

	"github.com/labstack/echo/v4"
)

func homeController(c echo.Context) error {
	return c.Render(200, "index", map[string]interface{}{
		"cpu_count":           runtime.NumCPU(),
		"cpu_freq":            600.,
		"cpu_mem_avail":       22.4,
		"cpu_mem_free":        0.11,
		"cpu_mem_total":       232,
		"cpu_mem_used":        0.43,
		"cpu_percent":         1.8,
		"disk_usage_free":     24.6,
		"disk_usage_percent":  17.7,
		"disk_usage_total":    31.3,
		"disk_usage_used":     52.9,
		"sensor_temperatures": 52.6,
	})
}

func routes(e *echo.Echo) {
	e.GET("/", homeController)
}
