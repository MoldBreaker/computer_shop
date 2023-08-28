package models

import (
	"time"
)

/*
CREATE TABLE Users (
    user_id INT NOT NULL AUTO_INCREMENT,
    role_id INT NOT NULL,
    user_name VARCHAR(255) NULL,
    email VARCHAR(255) NULL,
    password VARCHAR(255) NULL,
    avatar VARCHAR(255) NULL,
    token VARCHAR(255) NULL,
    phone VARCHAR(20) NULL,
    address VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);
*/

type UserModel struct {
	UserId    int        `json:"user_id" form:"user_id"`
	RoleId    int        `json:"role_id" form:"role_id"`
	Username  string     `json:"user_name" form:"user_name"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password" form:"password"`
	Avatar    string     `json:"avatar" form:"avatar"`
	Token     string     `json:"token" form:"token"`
	Phone     string     `json:"phone" form:"phone"`
	Address   string     `json:"address" form:"address"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"`
}

func (UM *UserModel) GetTableName() string {
	return `Users`
}
