package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"database/sql"
)

type InvoiceDetailDAO struct {
}

func (InvoiceDetailDAO *InvoiceDetailDAO) FindById(id int) (models.InvoiceDetailModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Invoice_Detail WHERE product_id = ? AND invoice_id = ?"
	var invoiceDetail models.InvoiceDetailModel
	if err := db.QueryRow(query, id).Scan(&invoiceDetail.ProductId, &invoiceDetail.InvoiceId, &invoiceDetail.Quantity, &invoiceDetail.ProductPrice); err != nil {
		return invoiceDetail, err
	}
	return invoiceDetail, nil
}

func (InvoiceDetailDAO *InvoiceDetailDAO) Create(invoiceDetail models.InvoiceDetailModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Invoice_Detail(product_id, invoice_id, quantity, product_price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, invoiceDetail.ProductId, invoiceDetail.InvoiceId, invoiceDetail.Quantity, invoiceDetail.ProductPrice)
	if err != nil {
		return err
	}
	return nil
}

func ScanToInvoiceDetail(rows *sql.Rows) (*models.InvoiceDetailModel, error) {
	var InvoiceDetail *models.InvoiceDetailModel
	if err := rows.Scan(&InvoiceDetail.ProductId, &InvoiceDetail.InvoiceId, &InvoiceDetail.Quantity, &InvoiceDetail.ProductPrice); err != nil {
		return nil, err
	}
	return InvoiceDetail, nil
}

func (InvoiceDetailDAO *InvoiceDetailDAO) FindAll() ([]models.InvoiceDetailModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Invoice_Detail"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var InvoiceDetails []models.InvoiceDetailModel
	for rows.Next() {
		invoiceDetail, err := ScanToInvoiceDetail(rows)
		if err != nil {
			return nil, err
		}
		InvoiceDetails = append(InvoiceDetails, *invoiceDetail)
	}
	return InvoiceDetails, nil
}

func (InvoiceDetailDAO *InvoiceDetailDAO) Update(invoiceDetail models.InvoiceDetailModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Invoice_Detail SET quantity = ?, product_price = ? WHERE invoice_id = ? AND product_id = ?"
	_, err := db.Exec(query, invoiceDetail.Quantity, invoiceDetail.ProductPrice, invoiceDetail.InvoiceId, invoiceDetail.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (InvoiceDetailDAO *InvoiceDetailDAO) Delete(invoiceId int, productId int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Invoice_Detail WHERE invoice_id = ? AND product_id = ?"
	_, err := db.Exec(query, invoiceId, productId)
	if err != nil {
		return err
	}
	return nil
}

func (InvoiceDetailDAO *InvoiceDetailDAO) FindByCondition(condition string) ([]models.InvoiceDetailModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Invoice_Detail " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var InvoiceDetails []models.InvoiceDetailModel
	for rows.Next() {
		var InvoiceDetail models.InvoiceDetailModel
		if err := rows.Scan(&InvoiceDetail.ProductId, &InvoiceDetail.InvoiceId, &InvoiceDetail.Quantity, &InvoiceDetail.ProductPrice); err != nil {
			return nil, err
		}
		InvoiceDetails = append(InvoiceDetails, InvoiceDetail)
	}
	return InvoiceDetails, nil
}
