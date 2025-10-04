package main

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed public/*
var embeddedFiles embed.FS

//go:embed views/*.html
var templateFiles embed.FS

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseFS(templateFiles, "views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	routes(e)

	publicFS, err := fs.Sub(embeddedFiles, "public")
	if err != nil {
		log.Fatal(err)
	}
	e.StaticFS("/", publicFS)

	e.Logger.Fatal(e.Start(":3000"))
}
