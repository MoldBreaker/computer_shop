package models

import (
	"database/sql"
	"time"
)

/*CREATE TABLE Notifications (
    notification_id INT NOT NULL AUTO_INCREMENT,
    content VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT NOT NULL,
    PRIMARY KEY (notification_id)
);*/

type NotificationModel struct {
	NotificationId int
	Content string
	CreatedAt *time.Time
	UserId int
}

func (NM *NotificationModel) GetTableName() string {
	return `Notifications`
}

func(a *NotificationModel) ScanToNotificationModel(rows *sql.Rows) (*NotificationModel, error) {
	err := rows.Scan(&a.NotificationId, &a.Content, &a.CreatedAt, a.UserId)
	if err != nil {
		return nil, err
	}
	return a, nil
}
