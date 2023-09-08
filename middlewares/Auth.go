package middlewares

import (
	"computer_shop/dao"
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthMiddleware struct {
}

func (AuthMiddleware *AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := helpers.GetCookie("remember", c)
		if cookie != "" {
			var UserDAO dao.UserDAO
			query := "WHERE token = '" + cookie + "'"
			result, _ := UserDAO.FindByCondition(query)
			if len(result) > 0 {
				helpers.SetSession("user", result[0], c)
			}
		}
		return next(c)
	}
}

func (AuthMiddleware *AuthMiddleware) IsLogined(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := helpers.GetSession("user", c)
		if user == nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Bạn cần phải đăng nhập để sử dụng chức năng này",
			})
		}
		return next(c)
	}
}

func (AuthMiddleware *AuthMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := helpers.GetSession("user", c)
		userModel := user.(models.UserModel)
		if userModel.RoleId != utils.Admin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Bạn không có quyền truy cập vào chức năng này",
			})
		}
		return next(c)
	}
}

func (AuthMiddleware *AuthMiddleware) IsSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := helpers.GetSession("user", c)
		userModel := user.(models.UserModel)
		if userModel.RoleId != utils.SuperAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Bạn không có quyền truy cập vào chức năng này",
			})
		}
		return next(c)
	}
}

func (AuthMiddleware *AuthMiddleware) IsAdminOrSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := helpers.GetSession("user", c)
		userModel := user.(models.UserModel)
		if userModel.RoleId != utils.Admin && userModel.RoleId != utils.SuperAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "Bạn không có quyền truy cập vào chức năng này",
			})
		}
		return next(c)
	}
}
