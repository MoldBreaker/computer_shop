package middlewares

import (
	"computer_shop/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthRedirect struct{}

func (AuthRedirect *AuthRedirect) IsLogined(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := helpers.GetSession("user", c)
		if user == nil {
			http.Redirect(c.Response(), c.Request(), "/auth", http.StatusSeeOther)
			return nil
		}
		return next(c)
	}
}
