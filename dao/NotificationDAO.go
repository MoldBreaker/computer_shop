package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type NotificationDAO struct {
}

func (NotificationDAO *NotificationDAO) Create(notification models.NotificationModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Notifications (content, user_id) values (?, ?)"
	result, err := db.Exec(query, notification.Content, notification.UserId)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)
}
