package controllers

import (
	"computer_shop/models"
	"computer_shop/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct{}

var UserService services.UserService

func (UserController *UserController) Register(e echo.Context) error {
	var user models.UserModel
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_, errStr, err := UserService.Register(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	return e.String(http.StatusOK, "Dang ki thnanh cong")
}

func (UserController *UserController) Login(e echo.Context) error {
	var user models.UserModel
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userResult, errStr, err := UserService.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	if errStr != "" {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}

	return e.JSON(http.StatusOK, userResult)
}
