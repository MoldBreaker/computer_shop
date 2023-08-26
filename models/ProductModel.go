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
	CategoryId  int    `json:"category_id" form:"category_id"`
	ProductName string `json:"product_name" form:"product_name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	IsDeleted   bool
}

type ProductModelRequest struct {
	CategoryId  int                 `json:"category_id" form:"category_id"`
	ProductName string              `json:"product_name" form:"product_name"`
	Description string              `json:"description" form:"description"`
	Price       int                 `json:"price" form:"price"`
	Images      []ProductImageModel `json:"images" form:"images"`
}

func (PM *ProductModel) GetTableName() string {
	return `Products`
}
