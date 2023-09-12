package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"computer_shop/utils"
	"errors"
	"fmt"
)

type NotificationService struct {
}

var (
	NotificationDAO dao.NotificationDAO
)

func (NotificationService *NotificationService) SendNotificationsToAllUser(message string) error {
	condition := fmt.Sprintf("WHERE role_id = %d", utils.User)
	users, err := UserDAO.FindByCondition(condition)
	if err != nil {
		return errors.New("không thể tìm thấy user trong hệ thống")
	}
	for _, user := range users {
		notification := models.NotificationModel{
			UserId:  user.UserId,
			Content: message,
		}
		NotificationDAO.Create(notification)
	}
	return nil
}

func (NotificationService *NotificationService) SendNotificationToOneUser(userId int, message string) error {
	notification := models.NotificationModel{
		UserId:  userId,
		Content: message,
	}
	var id int
	id = NotificationDAO.Create(notification)
	if id == 0 {
		return errors.New("có lỗi xảy ra khi gửi thông báo")
	}
	return nil
}

func (NotificationService *NotificationService) DeleteNotification(notificationId int) error {
	return NotificationDAO.Delete(notificationId)
}

func (NotificationService *NotificationService) GetAllNotification(userId int) ([]models.NotificationModel, error) {
	condition := fmt.Sprintf("WHERE user_id = %d ORDER BY created_at DESC", userId)
	return NotificationDAO.FindByCondition(condition)
}
