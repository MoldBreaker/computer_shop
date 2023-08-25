package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type RoleDAO struct {
}

func (RoleDAO *RoleDAO) Create(role models.RoleModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Role (role_name) values (?)"
	result, err := db.Exec(query, &role.RoleName)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)

}

func (RoleDAO *RoleDAO) FindAll() ([]models.RoleModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM role"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var Role []models.RoleModel
	for rows.Next() {
		var role models.RoleModel
		err := rows.Scan(&role.RoleId, &role.RoleName)
		if err != nil {
			return nil, err
		}
		Role = append(Role, role)
	}
	defer rows.Close()
	return Role, nil

}
