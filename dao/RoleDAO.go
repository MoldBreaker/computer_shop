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
	result, err := db.Exec(query, role.RoleName)
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

func (RoleDAO *RoleDAO) FindById(id int) (models.RoleModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM role WHERE role_id =?"
	row := db.QueryRow(query, id)
	var role models.RoleModel
	err := row.Scan(&role.RoleId, &role.RoleName)
	if err != nil {
		return role, err
	}
	return role, nil
}

func (RoleDAO *RoleDAO) Update(role models.RoleModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE role SET role_name =? WHERE role_id =?"
	_, err := db.Exec(query, role.RoleName, role.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (RoleDAO *RoleDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM role WHERE role_id =?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (RoleDAO *RoleDAO) FindByCondition(condition string) ([]models.RoleModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM role WHERE " + condition
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
