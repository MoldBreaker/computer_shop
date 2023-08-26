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
	ProductSrevice services.ProductService
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
	result := ProductSrevice.GetProductList(page, search, sort, col)
	if result == nil {
		return e.String(http.StatusBadRequest, "Bad Request")
	}
	return e.JSON(http.StatusOK, result)
}

func (ProductController *ProductController) CreateProduct(e echo.Context) error {
	var product models.ProductModel
	if err := e.Bind(&product); err != nil {
		return err
	}
	id := ProductSrevice.Create(product)
	product = ProductSrevice.GetProductById(id)
	if product.ProductId == 0 {
		return echo.NewHTTPError(500, "Error creating product")
	}
	return e.JSON(200, product)
}

func (ProductController *ProductController) GetProductById(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	product := ProductSrevice.GetProductById(id)
	if product.ProductId == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}
	return e.JSON(http.StatusOK, product)
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
	if !ProductSrevice.UpdateProduct(product) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating product")
	}
	return e.JSON(http.StatusOK, product)
}

func (ProductController *ProductController) DeleteProduct(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	if !ProductSrevice.DeleteProduct(id) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting product")
	}
	return e.NoContent(http.StatusOK)
}
