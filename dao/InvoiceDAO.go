package dao

import (
	"computer_shop/config"
	"computer_shop/models"
	"database/sql"
)

type InvoiceDAO struct {
}

/*
CREATE TABLE Invoices (
    invoice_id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    total_price INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (invoice_id)
);
*/

func (InvoiceDAO InvoiceDAO) Create(invoice models.InvoiceModel) int {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO invoices(user_id, total_price) VALUES(?, ?)"
	result, err := db.Exec(query, invoice.UserId, invoice.TotalPrice)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1
	}
	return int(id)
}

func (InvoiceDAO InvoiceDAO) FindAll() ([]models.InvoiceModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM invoices"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Invoices []models.InvoiceModel
	for rows.Next() {
		invoice, err := ScanToInvoiceModel(rows)
		if err != nil {
			return nil, err
		}
		Invoices = append(Invoices, *invoice)
	}
	return Invoices, nil
}

func (InvoiceDAO InvoiceDAO) FindById(id int) (models.InvoiceModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM invoices WHERE invoice_id = ?"
	var Invoice models.InvoiceModel
	if err := db.QueryRow(query, id).Scan(&Invoice.InvoiceId, &Invoice.UserId, &Invoice.TotalPrice, &Invoice.CreatedAt); err != nil {
		return Invoice, err
	}
	return Invoice, nil
}

func (InvoiceDAO InvoiceDAO) Update(invoice models.InvoiceModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE invoices SET user_id = ?, total_price = ? WHERE invoice_id = ?"
	_, err := db.Exec(query, invoice.UserId, invoice.TotalPrice, invoice.InvoiceId)
	if err != nil {
		return err
	}
	return nil
}

func (InvoiceDAO InvoiceDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM invoices WHERE invoice_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (InvoiceDAO InvoiceDAO) FindByCondition(condition string) ([]models.InvoiceModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM invoices " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var invoices []models.InvoiceModel
	for rows.Next() {
		var invoice models.InvoiceModel
		if err := rows.Scan(&invoice.InvoiceId, &invoice.UserId, &invoice.TotalPrice, &invoice.CreatedAt); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func ScanToInvoiceModel(rows *sql.Rows) (*models.InvoiceModel, error) {
	var invoiceModel *models.InvoiceModel
	err := rows.Scan(&invoiceModel.InvoiceId, &invoiceModel.UserId, &invoiceModel.TotalPrice, invoiceModel.CreatedAt)
	if err != nil {
		return nil, err
	}
	return invoiceModel, nil
}
