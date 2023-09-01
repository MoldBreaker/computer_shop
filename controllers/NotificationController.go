package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type NotificationController struct {
}

func (NotificationController *NotificationController) GetAllNotifications(e echo.Context) error {
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	notifications, err := NotificationService.GetAllNotification(user.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Lỗi khi lấy thông báo")
	}
	return e.JSON(http.StatusOK, notifications)
}

func (NotificationController *NotificationController) DelateNotification(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID không hợp lệ")
	}
	if err := NotificationService.DeleteNotification(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Lỗi khi xóa thông báo")
	}
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	if notifications, err := NotificationService.GetAllNotification(user.UserId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Lỗi khi lấy thông báo")
	} else {
		return e.JSON(http.StatusOK, notifications)
	}
}
