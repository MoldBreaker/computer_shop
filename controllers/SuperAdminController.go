package controllers

import (
	"computer_shop/helpers"
	"github.com/labstack/echo/v4"
	"html/template"
)

type SuperAdminController struct {
}

func (SuperAdminController *SuperAdminController) RenderSuperAminPage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	tmpl := template.Must(template.ParseFiles("views/superadmin/superadmin.html"))
	return tmpl.Execute(e.Response(), map[string]interface{}{
		"User": userSession,
	})
}
