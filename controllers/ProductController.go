package controllers

import (
	"computer_shop/models"
	"computer_shop/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
}

var (
	ProductImageService services.ProductImageService
	ProductService      services.ProductService
	NotificationService services.NotificationService
)

func (ProductController *ProductController) GetListProducts(e echo.Context) error {
	pageStr := e.QueryParam("page")
	var page int
	pageConv, err := strconv.Atoi(pageStr)
	if err == nil {
		page = pageConv
	} else {
		page = 1
	}
	search := e.QueryParam("search")
	sort := e.QueryParam("sort")
	col := e.QueryParam("col")
	result := ProductService.GetProductList(page, search, sort, col)
	if len(result) == 0 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "You have no results",
		})
	}
	var responseList []models.ProductResponse
	for i := 0; i < len(result); i++ {
		var response models.ProductResponse
		urls, _ := ProductImageService.GetImagesByProductId(result[i].ProductId)
		responseList = append(responseList, response.Parse(result[i], urls))
	}
	resLength := ProductService.GetProductsLength(search, sort, col)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"products":  responseList,
		"maxLength": resLength,
	})
}

func (ProductController *ProductController) CreateProduct(e echo.Context) error {
	form, err := e.MultipartForm()
	if err != nil {
		return e.String(http.StatusBadRequest, "Not a multipart form")
	}
	files := form.File["images"]
	var product models.ProductModel
	if err := e.Bind(&product); err != nil {
		return err
	}
	id := ProductService.Create(product)
	product = ProductService.GetProductById(id)
	if product.ProductId == 0 {
		return echo.NewHTTPError(500, "Error when create product")
	} else {
		result, err := ProductImageService.CreateMultipleImages(files, product.ProductId)
		if err != nil {
			return echo.NewHTTPError(500, result)
		}
	}
	var response models.ProductResponse
	urls, _ := ProductImageService.GetImagesByProductId(product.ProductId)

	//Send notifiactions
	if err := NotificationService.SendNotificationsToAllUser(product.ProductName + " đã được thêm vào cửa hàng"); err != nil {
		return echo.NewHTTPError(500, err)
	}

	return e.JSON(200, response.Parse(product, urls))
}

func (ProductController *ProductController) GetProductById(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	product := ProductService.GetProductById(id)
	if product.ProductId == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}
	var response models.ProductResponse
	urls, _ := ProductImageService.GetImagesByProductId(product.ProductId)
	return e.JSON(200, response.Parse(product, urls))
}

func (ProductController *ProductController) UpdateProduct(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	var product models.ProductModel
	if err := e.Bind(&product); err != nil {
		return err
	}
	product.ProductId = id
	if !ProductService.UpdateProduct(product) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating product")
	}
	var response models.ProductResponse
	urls, _ := ProductImageService.GetImagesByProductId(product.ProductId)
	return e.JSON(200, response.Parse(product, urls))
}

func (ProductController *ProductController) DeleteProduct(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	if !ProductService.DeleteProduct(id) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting product")
	}
	return e.NoContent(http.StatusOK)
}
