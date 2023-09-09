package services

import (
	"computer_shop/dao"
	"computer_shop/models"
)

type CategoryService struct{}

var (
	CategoryDAO dao.CategoryDAO
)

func (c *CategoryService) GetAllCategory() ([]models.CategoryModel, error) {
	return CategoryDAO.FindAll()
}
