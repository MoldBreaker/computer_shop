package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
)

type RoleService struct {
}

var (
	RoleDAO dao.RoleDAO
)

func (RoleService *RoleService) UpdateRole(role models.RoleModel) int {
	id := RoleDAO.Create(role)
	return id
}

func (RoleService *RoleService) GetRoleById(id int) models.RoleModel {
	role, _ := RoleDAO.FindById(id)
	return role
}
func (RoleService *RoleService) GetAllRoles() ([]models.RoleModel, error) {
	return RoleDAO.FindAll()
}

func (RoleService *RoleService) UpdateUserRole(userId, roleId int) error {
	user, err := UserDAO.FindById(userId)
	if err != nil {
		return errors.New("Error when getting user")
	}
	user.RoleId = roleId
	return UserDAO.Update(user)
}
