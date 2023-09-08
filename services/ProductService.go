package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"fmt"
	"os"
	"strconv"
	"time"
)

type ProductService struct {
}

var (
	ProductDAO dao.ProductDAO
)

func GenerateQuery(itemPerPage int, page int, search, sort, col string) string {

	query := ""

	// Search condition
	if search != "" {
		query += fmt.Sprintf(" WHERE is_deleted = false AND product_name LIKE '%%%s%%'", search)
	} else {
		query += fmt.Sprintf(" WHERE is_deleted = false")
	}

	// Sorting
	if sort != "" && col != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", col, sort)
	}

	// Pagination
	offset := (page - 1) * itemPerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", itemPerPage, offset)

	return query
}

func GenerateQueryNotPaging(search, sort, col string) string {

	query := ""

	// Search condition
	if search != "" {
		query += fmt.Sprintf(" WHERE is_deleted = false AND product_name LIKE '%%%s%%'", search)
	} else {
		query += fmt.Sprintf(" WHERE is_deleted = false")
	}

	// Sorting
	if sort != "" && col != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", col, sort)
	}

	return query
}

func (ProductService *ProductService) GetProductList(page int, search string, sort string, col string) []models.ProductModel {
	itemPerPageStr := os.Getenv("ITEM_PER_PAGE")
	itemPerPage, _ := strconv.Atoi(itemPerPageStr)
	result, err := ProductDAO.FindByCondition(GenerateQuery(itemPerPage, page, search, sort, col))
	if err != nil {
		return result
	}
	return result
}

func (ProductService *ProductService) GetProductsLength(search string, sort string, col string) int {
	result, err := ProductDAO.FindByCondition(GenerateQueryNotPaging(search, sort, col))
	if err != nil {
		return -1
	}
	return len(result)
}

func (ProductService *ProductService) Create(product models.ProductModel) int {
	return ProductDAO.Create(product)
}

func (ProductService *ProductService) GetProductById(id int) models.ProductModel {
	result, err := ProductDAO.FindById(id)
	if err != nil {
		return models.ProductModel{}
	}
	return result
}

func (ProductService *ProductService) UpdateProduct(product models.ProductModel) bool {
	return ProductDAO.Update(product) == nil
}

func (ProductService *ProductService) DeleteProduct(id int) bool {
	product := ProductService.GetProductById(id)
	product.IsDeleted = true
	currentTime := time.Now()
	var currentTimePtr *time.Time
	currentTimePtr = &currentTime
	product.DeletedAt = currentTimePtr
	return ProductService.UpdateProduct(product)
}
