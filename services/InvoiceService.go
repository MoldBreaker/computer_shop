package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
	"fmt"
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

func (InvoiceService *InvoiceService) GetHistoryInvoices(user models.UserModel) ([]models.InvoiceModel, error) {
	query := fmt.Sprintf("WHERE user_id = %d", user.UserId)
	return InvoiceDAO.FindByCondition(query)
}

func (InvoiceService *InvoiceService) GetInvoiceDetails(invoice_id int) (models.InvoiceModel, error) {
	return InvoiceDAO.FindById(invoice_id)
}
