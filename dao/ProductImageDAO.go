package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"log"
)

type ProductImageDAO struct {
}

func (ProductImageDAO *ProductImageDAO) Create(productImage models.ProductImageModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Product_Images(productid, link) values(?,?)"
	result, err := db.Exec(query, productImage.ImageId, productImage.ProductId, productImage.Link)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id)
}

func (ProductImageDAO *ProductImageDAO) FindAll() ([]models.ProductImageModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Product_Images"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var ProductImages []models.ProductImageModel
	for rows.Next() {
		var productImage models.ProductImageModel
		rows.Scan(&productImage.ImageId, productImage.ProductId, productImage.Link)
		ProductImages = append(ProductImages, productImage)
	}
	rows.Close()
	return ProductImages, nil
}
