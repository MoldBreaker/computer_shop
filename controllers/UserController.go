package controllers

import (
	"computer_shop/helpers"
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
	var validator helpers.Validator
	validator.Chain = append(validator.Chain, validator.Required(user.Username, "username không được để trống"))
	validator.Chain = append(validator.Chain, validator.Required(user.Email, "email không được để trống"))
	validator.Chain = append(validator.Chain, validator.IsEmail(user.Email, "email không hợp lệ"))
	validator.Chain = append(validator.Chain, validator.Required(user.Password, "mật khẩu không được để trống"))
	validator.Chain = append(validator.Chain, validator.MinLength(user.Password, 6, "mật khẩu phái có ít nhất 6 kí tự"))
	if err := validator.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	_, errStr, err := UserService.Register(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Đăng kí thành công",
	})
}

func (UserController *UserController) Login(e echo.Context) error {
	var user models.UserModel
	if err := e.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var validatorLogin helpers.Validator
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.Required(user.Email, "email không được để trống"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.IsEmail(user.Email, "email không hợp lệ"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.Required(user.Password, "mật khẩu không được để trống"))
	validatorLogin.Chain = append(validatorLogin.Chain, validatorLogin.MinLength(user.Password, 6, "mật khẩu phái có ít nhất 6 kí tự"))
	if err := validatorLogin.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userResult, errStr, err := UserService.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	if errStr != "" {
		return echo.NewHTTPError(http.StatusBadRequest, errStr)
	}
	errSession := helpers.SetSession("user", userResult, e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errSession)
	}
	token := helpers.GenarateToken()
	errSetToken := UserService.SetToken(userResult, token)
	if errSetToken != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errSetToken)
	}
	helpers.SetCookie("remember", token, e)
	return e.JSON(http.StatusOK, userResult)
}

func (UserController *UserController) Logout(e echo.Context) error {
	err := helpers.RemoveSession("user", e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Bạn chưa đăng nhập")
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Đăng xuất thành công",
	})
}

func (UserController *UserController) ResetPassword(e echo.Context) {

}
