package controllers

import (
	"computer_shop/helpers"
	"computer_shop/models"
	"computer_shop/services"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type InvoiceController struct {
}

var (
	InvoiceService       services.InvoiceService
	InvoiceDetailService services.InvoiceDetailService
)

func (InvoiceController *InvoiceController) CreateInvoice(e echo.Context) error {
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	data := map[string]interface{}{
		"cart": [][]int{},
	}
	err := json.NewDecoder(e.Request().Body).Decode(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "không thể chuyển đổi qua json")
	}

	cartInterface, ok := data["cart"].([]interface{})
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "không thể chuyển đổi qua dữ liệu")
	}

	cartData := make([][]int, len(cartInterface))
	for i, item := range cartInterface {
		pair, ok := item.([]interface{})
		if !ok || len(pair) != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, "dữ liệu không hơ lệ")
		}
		first, ok := pair[0].(float64) // Assuming the values are integers in the JSON
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "dữ liệu không hơp lệ")
		}
		second, ok := pair[1].(float64)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "dữ liệu không hơp lệ")
		}
		cartData[i] = []int{int(first), int(second)}
	}

	errInsert := InvoiceService.Create(user, cartData)
	if errInsert != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errInsert)
	}
	return e.JSON(http.StatusOK, "Thanh toán thành công")
}

func (InvoiceController *InvoiceController) GetHistoryInvoices(e echo.Context) error {
	userModel, errSession := helpers.GetSession("user", e)
	if errSession != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bạn chưa đăng nhập")
	}
	user := userModel.(models.UserModel)
	invoices, err := InvoiceService.GetHistoryInvoices(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return e.JSON(http.StatusOK, invoices)
}

func (InvoiceController *InvoiceController) GetHistoryInvoiceDetails(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "mã hóa đơn khng hợp lệ")
	}
	invoice, err := InvoiceService.GetInvoiceDetails(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	details, err := InvoiceDetailService.GetHistoryInvoiceDetails(id)
	var invoiceResponse []models.InvoiceDetailResponseModel
	for _, detail := range details {
		product := ProductService.GetProductById(detail.ProductId)
		var productResponse models.ProductResponse
		urls, err := ProductImageService.GetImagesByProductId(product.ProductId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "không thể lấy ảnh sản phẩm")
		}
		invoiceResponse = append(invoiceResponse, models.InvoiceDetailResponseModel{
			Product:      productResponse.Parse(product, urls),
			Quantity:     detail.Quantity,
			ProductPrice: detail.ProductPrice,
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"invoice": invoice,
		"details": invoiceResponse,
	})
}
