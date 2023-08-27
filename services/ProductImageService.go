package services

import (
	"computer_shop/dao"
	"computer_shop/helpers"
	"computer_shop/models"
	"mime/multipart"
	"strconv"
)

var (
	ProductImageDAO dao.ProductImageDAO
)

type ProductImageService struct {
}

func (p *ProductImageService) CreateMultipleImages(files []*multipart.FileHeader, productId int) (string, error) {
	URLs, errStr, err := helpers.UploadFiles(files)
	if err != nil {
		return errStr, err
	}

	for i := 0; i < len(URLs); i++ {
		productImage := models.ProductImageModel{
			ProductId: productId,
			Link:      URLs[i],
		}
		ProductImageDAO.Create(productImage)
	}

	return strconv.Itoa(len(URLs)) + " file(s) uploaded", nil
}

func (p *ProductImageService) GetImagesByProductId(productId int) ([]string, error) {
	condition := "WHERE product_id = " + strconv.Itoa(productId)
	results, err := ProductImageDAO.FindByCondition(condition)
	if err != nil {
		return nil, err
	}
	var URLs []string
	for i := 0; i < len(results); i++ {
		URLs = append(URLs, results[i].Link)
	}
	return URLs, nil
}
