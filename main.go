package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "public")

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", map[string]interface{}{
			"cpu_count":           4,
			"cpu_freq":            "[ 600.0, 600.0, 1200.0 ]",
			"cpu_mem_avail":       463953920,
			"cpu_mem_free":        115789824,
			"cpu_mem_total":       971063296,
			"cpu_mem_used":        436252672,
			"cpu_percent":         1.8,
			"disk_usage_free":     24678121472,
			"disk_usage_percent":  17.7,
			"disk_usage_total":    31307206656,
			"disk_usage_used":     5292728320,
			"sensor_temperatures": 52.616,
		})
	})
	e.Logger.Fatal(e.Start(":3000"))
}
