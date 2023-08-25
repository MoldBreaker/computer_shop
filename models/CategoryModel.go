package models

import (
	"time"
)

/*
CREATE TABLE Categories (
    category_id INT NOT NULL AUTO_INCREMENT,
    category_name VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (category_id)
);
*/

type CategoryModel struct {
	CategoryId   int
	CategoryName string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (CategoryModel *CategoryModel) GetTableName() string {
	return `Categories`
}
