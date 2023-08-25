package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitWebRoutes() *echo.Echo {
	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return router
}
