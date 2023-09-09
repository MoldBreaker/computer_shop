package controllers

import (
	"computer_shop/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct{}

var (
	CategoryService services.CategoryService
)

func (CategoryController *CategoryController) GetAllCategory(e echo.Context) error {
	categories, err := CategoryService.GetAllCategory()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not get category",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"categories": categories,
	})
}
