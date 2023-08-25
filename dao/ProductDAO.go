package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type ProductDAO struct {
}

func (ProductDAO *ProductDAO) Create(product models.ProductModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Products (category_id, productname, description, price) values (?,?,?,?)"
	result, err := db.Exec(query, product.ProductId, product.CategoryId, product.ProductName, product.Description, product.Price, product.CreatedAt, product.UpdatedAt, product.DeletedAt, product.IsDeleted)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)
}

func (ProductDAO *ProductDAO) FindAll() ([]models.ProductModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var Products []models.ProductModel
	for rows.Next() {
		var product models.ProductModel
		rows.Scan(&product.ProductId, &product.CategoryId, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt, &product.IsDeleted)
		Products = append(Products, product)
	}
	defer rows.Close()
	return Products, nil
}
