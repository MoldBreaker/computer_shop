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

func (NotificationDAO *NotificationDAO) FindALl() ([]models.NotificationModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Notifications"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var Notifications []models.NotificationModel
	for rows.Next() {
		var notification models.NotificationModel
		rows.Scan(&notification.NotificationId, &notification.Content, &notification.UserId)
		Notifications = append(Notifications, notification)
	}
	rows.Close()
	return Notifications, nil
}

func (NotificationDAO *NotificationDAO) Update(notification models.NotificationModel) bool {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Notifications SET content = ?, user_id =? WHERE notification_id =?"
	_, err := db.Exec(query, notification.Content, notification.UserId, notification.NotificationId)
	if err != nil {
		return false
	}
	return true
}

func (NotificationDAO *NotificationDAO) Delete(id int) bool {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Notifications WHERE NotificationId = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return false
	}
	return true
}
