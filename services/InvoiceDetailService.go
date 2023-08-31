package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
	"fmt"
)

type InvoiceDetailService struct {
}

var (
	InvoiceDetailDAO dao.InvoiceDetailDAO
)

func (InvoiceDetailService *InvoiceDetailService) Create(invoiceId, productId, quantity int) (int, error) {
	productModel, err := ProductDAO.FindById(productId)
	if err != nil {
		return 0, errors.New("không thể lấy được thông tin sản phẩm")
	}
	price := productModel.Price * quantity
	invoiceDetail := models.InvoiceDetailModel{
		ProductId:    productId,
		InvoiceId:    invoiceId,
		Quantity:     quantity,
		ProductPrice: price,
	}
	errCreate := InvoiceDetailDAO.Create(invoiceDetail)
	if errCreate != nil {
		return 0, errors.New("không thể lưu sản phẩm trong chi tiết hóa đơn")
	}
	return price, nil
}

func (InvoiceDetailService *InvoiceDetailService) GetHistoryInvoiceDetails(invoice_id int) ([]models.InvoiceDetailModel, error) {
	query := fmt.Sprintf("WHERE invoice_id = %d", invoice_id)
	return InvoiceDetailDAO.FindByCondition(query)
}
