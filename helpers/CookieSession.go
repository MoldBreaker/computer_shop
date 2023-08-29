package helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SetCookie(key, value string, e echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * 5 * time.Hour)
	cookie.Path = "/"
	e.SetCookie(cookie)
}

func GetCookie(key string, e echo.Context) (string, error) {
	cookie, err := e.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}
