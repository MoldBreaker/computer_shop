package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type UserDAO struct {
}

func (UserDAO *UserDAO) Create(user models.UserModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Users (role_id, user_name, email, password, avatar, token, phone, address) values (?,?,?,?,?,?,?,?)"
	result, err := db.Exec(query, user.RoleId, user.Username, user.Email, user.Password, user.Avatar, user.Token, user.Phone, user.Address)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)

}

func (UserDAO *UserDAO) FindAll() ([]models.UserModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var Users []models.UserModel
	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.UserId, &user.RoleId, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Token, &user.Phone, &user.Address, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		Users = append(Users, user)
	}
	defer rows.Close()
	return Users, nil
}
