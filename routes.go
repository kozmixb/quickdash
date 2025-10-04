package main

import (
	"github.com/labstack/echo/v4"
)

func homeController(c echo.Context) error {
	return c.Render(200, "index", map[string]interface{}{
		"lastlogin": "test",
		"datetime":  "asda",
	})
}

func routes(e *echo.Echo) {
	e.GET("/", homeController)
}
