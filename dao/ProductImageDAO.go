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
	query := "INSERT INTO Product_Images(product_id, link) values(?,?)"
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
		rows.Scan(&productImage.ImageId, &productImage.ProductId, &productImage.Link)
		ProductImages = append(ProductImages, productImage)
	}
	rows.Close()
	return ProductImages, nil
}

func (ProductImageDAO *ProductImageDAO) FindById(id int) (models.ProductImageModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Product_Images WHERE image_id = ?"
	var productImage models.ProductImageModel
	err := db.QueryRow(query, id).Scan(&productImage.ImageId, &productImage.ProductId, &productImage.Link)
	if err != nil {
		return productImage, err
	}
	return productImage, nil
}

func (ProductImageDAO *ProductImageDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Product_Images WHERE image_id =?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (ProductImageDAO *ProductImageDAO) Update(productImage models.ProductImageModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Product_Images SET product_id =?, link =? WHERE image_id =?"
	_, err := db.Exec(query, productImage.ProductId, productImage.Link, productImage.ImageId)
	if err != nil {
		return err
	}
	return nil
}

func (ProductImageDAO *ProductImageDAO) FindByCondition(condition string) ([]models.ProductImageModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Product_Images WHERE " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var ProductImages []models.ProductImageModel
	for rows.Next() {
		var productImage models.ProductImageModel
		rows.Scan(&productImage.ImageId, &productImage.ProductId, &productImage.Link)
		ProductImages = append(ProductImages, productImage)
	}
	rows.Close()
	return ProductImages, nil
}
