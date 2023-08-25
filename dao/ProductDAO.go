package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type ProductDAO struct {
}

/*
CREATE TABLE Products (
    product_id INT NOT NULL AUTO_INCREMENT,
    category_id INT NOT NULL,
    product_name VARCHAR(255) NULL,
    description TEXT NULL,
    price INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    is_deleted BOOLEAN NULL,
    PRIMARY KEY (product_id)
);
*/

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

func (ProductDAO *ProductDAO) FindById(id int) (models.ProductModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Products WHERE product_id = ?"
	var product models.ProductModel
	err := db.QueryRow(query, id).Scan(&product.ProductId, &product.CategoryId, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt, &product.IsDeleted)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (ProductDAO *ProductDAO) Update(product models.ProductModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Products SET category_id = ?, product_name = ?, description = ?, price = ?, deleted_at = ?, is_deleted = ? WHERE product_id = ?"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (ProductDAO *ProductDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Products WHERE product_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (ProductDAO *ProductDAO) FindByCondition(condition string) ([]models.ProductModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Products WHERE " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Products []models.ProductModel
	for rows.Next() {
		var product models.ProductModel
		rows.Scan(&product.ProductId, &product.CategoryId, &product.ProductName, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt, &product.IsDeleted)
		Products = append(Products, product)
	}
	return Products, nil
}
