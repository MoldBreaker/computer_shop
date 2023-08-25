package models

import (
	"database/sql"
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
	UserId    int
	RoleId    int
	Username  string
	Email     string
	Password  string
	Avatar    string
	Token     string
	Phone     string
	Address   string
	CreatedAt *time.Time
}

func (UM *UserModel) GetTableName() string {
	return `Users`
}

func(a *UserModel) ScanToUserModel(rows *sql.Rows) (*UserModel, error) {
	err := rows.Scan(&a.UserId, &a.RoleId, &a.Username, &a.Email, &a.Password, &a.Avatar, &a.Token, &a.Phone, &a.Address, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return a, nil
}