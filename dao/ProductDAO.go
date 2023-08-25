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
	result, err := db.Exec(query, product.ProductId, product.CategoryId, product.ProductName, product.Description, product.Price)
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

func (ProductDAO *ProductDAO) Update(product models.ProductModel) bool {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Products SET category_id = ?, product_name =?, description =?, price = ? WHERE product_id =?"
	_, err := db.Exec(query, product.CategoryId, product.ProductName, product.Description, product.Price, product.ProductId)
	if err != nil {
		return false
	}
	return true
}

func (ProductDAO *ProductDAO) Delete(id int) bool {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Products WHERE product_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return false
	}
	return true
}
