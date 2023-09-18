package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type NotificationDAO struct {
}

/*CREATE TABLE Notifications (
    notification_id INT NOT NULL AUTO_INCREMENT,
    content VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT NOT NULL,
    PRIMARY KEY (notification_id)
);*/

func (NotificationDAO NotificationDAO) Create(notification models.NotificationModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO notifications (content, user_id) values (?, ?)"
	result, err := db.Exec(query, notification.Content, notification.UserId)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)
}

func (NotificationDAO NotificationDAO) FindAll() ([]models.NotificationModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM notifications"
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var notifications []models.NotificationModel
	for rows.Next() {
		var notification models.NotificationModel
		err := rows.Scan(&notification.NotificationId, &notification.Content, &notification.CreatedAt, &notification.UserId)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (NotificationDAO NotificationDAO) FindById(id int) (models.NotificationModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM notifications WHERE notification_id = ?"
	var notification models.NotificationModel
	err := db.QueryRow(query, id).Scan(&notification.NotificationId, &notification.Content, &notification.CreatedAt, notification.UserId)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (NotificationDAO NotificationDAO) Update(notification models.NotificationModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE notifications SET content = ? WHERE notification_id = ?"
	_, err := db.Exec(query, notification.Content, notification.NotificationId)
	if err != nil {
		return err
	}
	return nil
}

func (NotificationDAO NotificationDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM notifications WHERE notification_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (NotificationDAO NotificationDAO) FindByCondition(condition string) ([]models.NotificationModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM notifications " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var notifications []models.NotificationModel
	for rows.Next() {
		var notification models.NotificationModel
		err := rows.Scan(&notification.NotificationId, &notification.Content, &notification.CreatedAt, &notification.UserId)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
