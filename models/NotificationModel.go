package models

import (
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
	Content        string
	CreatedAt      *time.Time
	UserId         int
}

func (NM *NotificationModel) GetTableName() string {
	return `Notifications`
}
