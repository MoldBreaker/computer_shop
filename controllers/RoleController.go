package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (RoleController *RoleController) UpdateUserRole(e echo.Context) error {
	userIdStr := e.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user id",
		})
	}
	var roleModel models.RoleModel
	if err := e.Bind(&roleModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid role code")
	}
	errResult := RoleService.UpdateUserRole(userId, roleModel.RoleId)
	if errResult != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errResult.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated user role successfully",
	})
}

func (RoleController *RoleController) GettAllRoles(e echo.Context) error {
	roles, err := RoleService.GetAllRoles()
	if err != nil {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"roles": roles,
	})
}
