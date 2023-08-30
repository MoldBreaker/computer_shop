package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleController struct{}

var (
	RoleService services.RoleService
)

func (RoleController *RoleController) CreateRole(e echo.Context) error {
	var roleModel models.RoleModel
	if err := e.Bind(&roleModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "không hợp lệ")
	}
	var validator helpers.Validator
	validator.Chain = append(validator.Chain, validator.Required(roleModel.RoleName, "role không được để trống"))
	if err := validator.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id := RoleService.UpdateRole(roleModel)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"role": RoleService.GetRoleById(id),
	})
}
