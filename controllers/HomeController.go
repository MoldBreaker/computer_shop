package controllers

import (
	"computer_shop/helpers"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func (HomeController *HomeController) RenderHomePage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	tmpl := template.Must(template.ParseFiles("views/home.html"))
	return tmpl.Execute(e.Response(), map[string]interface{}{
		"User": userSession,
	})
}

func (HomeController *HomeController) RenderAuthPage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	if userSession != nil {
		http.Redirect(e.Response(), e.Request(), e.Request().Header.Get("Referer"), 302)
	}
	tmpl := template.Must(template.ParseFiles("views/Auth.html"))
	return tmpl.Execute(e.Response(), nil)
}

func (HomeController *HomeController) RenderProductDetailPage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	tmpl := template.Must(template.ParseFiles("views/productDetail.html"))
	return tmpl.Execute(e.Response(), map[string]interface{}{
		"User": userSession,
	})
}
