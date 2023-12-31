package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController struct {
}

var CartService services.CartService

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

func (CartController *CartController) GetItemsInCart(e echo.Context) error {
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not logged in")
	}
	user := userModel.(models.UserModel)
	result, err := CartService.GetCartByUserId(user.UserId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, result)
}
