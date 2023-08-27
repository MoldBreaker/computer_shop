package models

import (
	"time"
)

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

type ProductModel struct {
	ProductId   int
	CategoryId  int        `json:"category_id" form:"category_id"`
	ProductName string     `json:"product_name" form:"product_name"`
	Description string     `json:"description" form:"description"`
	Price       int        `json:"price" form:"price"`
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`
	IsDeleted   bool       `json:"is_deleted" form:"is_deleted"`
}

type ProductModelRequest struct {
	CategoryId  int                 `json:"category_id" form:"category_id"`
	ProductName string              `json:"product_name" form:"product_name"`
	Description string              `json:"description" form:"description"`
	Price       int                 `json:"price" form:"price"`
	Images      []ProductImageModel `json:"images" form:"images"`
}

type ProductResponse struct {
	ProductId   int        `json:"product_id" form:"product_id"`
	CategoryId  int        `json:"category_id" form:"category_id"`
	ProductName string     `json:"product_name" form:"product_name"`
	Description string     `json:"description" form:"description"`
	Price       int        `json:"price" form:"price"`
	CreatedAt   *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" form:"deleted_at"`
	IsDeleted   bool       `json:"is_deleted" form:"is_deleted"`
	Images      []string   `json:"images" form:"images"`
}

func (PM *ProductModel) GetTableName() string {
	return `Products`
}

func (ProductResponse *ProductResponse) Parse(product ProductModel, images []string) ProductResponse {
	ProductResponse.ProductId = product.ProductId
	ProductResponse.CategoryId = product.CategoryId
	ProductResponse.ProductName = product.ProductName
	ProductResponse.Description = product.Description
	ProductResponse.Price = product.Price
	ProductResponse.CreatedAt = product.CreatedAt
	ProductResponse.UpdatedAt = product.UpdatedAt
	ProductResponse.DeletedAt = product.DeletedAt
	ProductResponse.IsDeleted = product.IsDeleted
	ProductResponse.Images = images
	return *ProductResponse
}
