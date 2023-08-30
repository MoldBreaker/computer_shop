package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
)

type InvoiceService struct {
}

var (
	InvoiceDAO dao.InvoiceDAO
)

func (InvoiceService *InvoiceService) Create(userModel models.UserModel, cartData [][]int) error {
	var InvoiceDetailService InvoiceDetailService
	invoice := models.InvoiceModel{
		UserId:     userModel.UserId,
		TotalPrice: 0,
	}
	id := InvoiceDAO.Create(invoice)

	sum := 0
	for _, pair := range cartData {
		price, err := InvoiceDetailService.Create(id, pair[0], pair[1])
		if err != nil {
			return errors.New("lỗi khi nhập vào chi tiết")
		}
		sum += price
	}
	invoiceResult := models.InvoiceModel{
		InvoiceId:  id,
		UserId:     userModel.UserId,
		TotalPrice: sum,
	}
	return InvoiceDAO.Update(invoiceResult)
}
