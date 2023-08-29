package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CartController struct {
}

var CartService services.CartServive

func (CartController *CartController) AddToCart(e echo.Context) error {
	productIdStr := e.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	userModel, err := helpers.GetSession("user", e)
	if userModel == nil || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Mày chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	userId := user.UserId
	errAdd := CartService.AddToCart(userId, productId)
	if errAdd != nil {
		return e.JSON(http.StatusBadRequest, errAdd)
	}
	return e.JSON(http.StatusOK, "Thêm sản phẩm vào giỏ hàng thành công")
}

func (CartController *CartController) UpdateInCart(e echo.Context) error {
	productIdtr := e.Param("id")
	productId, err := strconv.Atoi(productIdtr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Id không tồn tại")
	}
	types := e.QueryParams().Get("type")
	userModel, err := helpers.GetSession("user", e)
	if userModel == nil || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Mày chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	userId := user.UserId
	errUpdate := CartService.UpdateInCart(userId, productId, types)
	if errUpdate != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errUpdate)
	}
	return e.JSON(http.StatusOK, "Cập nhật thành công")
}

func (CartController *CartController) DeleteInCart(e echo.Context) error {
	productIdtr := e.Param("id")
	productId, err := strconv.Atoi(productIdtr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Id không tồn tại")
	}
	userModel, err := helpers.GetSession("user", e)
	if userModel == nil || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Mày chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	userId := user.UserId
	errDelete := CartService.DeleteInCart(userId, productId)
	if errDelete != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errDelete)
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Xóa sản phẩm thành công",
	})
}
