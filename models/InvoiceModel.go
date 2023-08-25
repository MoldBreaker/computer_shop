package models

import (
	"database/sql"
	"time"
)

/*
CREATE TABLE Invoices (
    invoice_id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    total_price INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (invoice_id)
);
*/

type InvoiceModel struct {
	InvoiceId int
	UserId int
	TotalPrice int
	CreatedAt *time.Time
}

func(IM *InvoiceModel) GetTableName() string {
	return `Invoices`
}

func(IM *InvoiceModel) ScanToInvoiceModel(rows *sql.Rows) (*InvoiceModel, error) {
	err := rows.Scan(&IM.InvoiceId, &IM.UserId, &IM.TotalPrice, IM.CreatedAt)
	if err != nil {
		return nil, err
	}
	return IM, nil
}
