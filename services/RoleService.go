package services

import (
	"computer_shop/dao"
	"computer_shop/models"
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
