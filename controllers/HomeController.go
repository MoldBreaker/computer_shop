package controllers

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

type HomeController struct{}

func (HomeController *HomeController) RenderHomePage(e echo.Context) error {
	tmpl := template.Must(template.ParseFiles("views/home.html"))
	return tmpl.Execute(e.Response(), nil)
}
