package middlewares

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CheckCartMiddleware struct{}

var (
	CartService services.CartService
)

func (CheckCartMiddleware *CheckCartMiddleware) CheckCartEmpty(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userSession, _ := helpers.GetSession("user", c)
		user := userSession.(models.UserModel)
		cartItems, err := CartService.GetCartByUserId(user.UserId)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "Cart not found",
			})
		} else if len(cartItems) == 0 {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "Your cart is empty",
			})
		} else {
			return next(c)
		}
	}
}
