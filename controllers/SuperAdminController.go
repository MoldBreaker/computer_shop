package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/utils"
	"html/template"

	"github.com/labstack/echo/v4"
)

type SuperAdminController struct {
}

func (SuperAdminController *SuperAdminController) RenderSuperAminPage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	user := userSession.(models.UserModel)
	tmpl, _ := template.ParseFiles("views/template/adminTemplate.html", "views/superadmin/superadmin.html")
	return tmpl.Execute(e.Response(), map[string]interface{}{
		"User":         user,
		"IsSuperAdmin": user.RoleId == utils.SuperAdmin,
	})
}

func (SuperAdminController *SuperAdminController) RenderAdminPage(e echo.Context) error {
	userSession, _ := helpers.GetSession("user", e)
	user := userSession.(models.UserModel)
	tmpl, _ := template.ParseFiles("views/template/adminTemplate.html", "views/superadmin/admin.html")
	return tmpl.Execute(e.Response(), map[string]interface{}{
		"User":         user,
		"IsSuperAdmin": user.RoleId == utils.SuperAdmin,
	})
}
