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
	CategoryId  int
	ProductName string
	Description string
	Price       int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	IsDeleted   bool
}

func (PM *ProductModel) GetTableName() string {
	return `Products`
}
