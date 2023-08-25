package models

import "database/sql"

/*
CREATE TABLE Product_Images (
    image_id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    link VARCHAR(255) NULL,
    PRIMARY KEY (image_id)
);
*/

type ProductImageModel struct {
	ImageId int
	ProductId int
	Link string
}

func (PIM *ProductImageModel) GetTableName () string {
	return `Product_Images`
}

func(a *ProductImageModel) ScanToProductImageModel(rows *sql.Rows) (*ProductImageModel, error) {
	err := rows.Scan(&a.ImageId, &a.ProductId, &a.Link)
	if err != nil {
		return nil, err
	}
	return a, nil
}