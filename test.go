package main

import (
	"computer_shop/config"
	"computer_shop/dao"
	"computer_shop/models"
	"fmt"
)

func main() {
	config.LoadENV()
	var (
		UserDAO dao.UserDAO
	)
	user := models.UserModel{
		Username: "admin",
		Password: "123456",
		Email:    "tayhoang64@gmail.com",
		Avatar:   "https://",
		Token:    "12132",
		Phone:    "123456",
		Address:  "123456",
		RoleId:   5,
	}
	UserDAO.Create(user)

	fmt.Print(UserDAO.FindAll())

}
