package models

import (
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
	InvoiceId  int
	UserId     int
	TotalPrice int
	CreatedAt  *time.Time
}

func (IM *InvoiceModel) GetTableName() string {
	return `Invoices`
}
