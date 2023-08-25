package models

import "database/sql"

/*
CREATE TABLE Invoice_Detail (
    product_id INT NOT NULL,
    invoice_id INT NOT NULL,
    quantity INT NULL,
    product_price INT NULL,
    PRIMARY KEY (product_id, invoice_id)
);
*/

type InvoiceDetailModel struct {
	ProductId int
	InvoiceId int
	Quantity int
	ProductPrice int
}

func (InvoiceDetailModel *InvoiceDetailModel) GetTableName() string{
	return `Invoice_Detail`
}

func (InvoiceDetailModel *InvoiceDetailModel) ScanToInvoiceDetailModel(rows *sql.Rows) (*InvoiceDetailModel, error) {
	err := rows.Scan(&InvoiceDetailModel.ProductId, &InvoiceDetailModel.InvoiceId, &InvoiceDetailModel.Quantity, &InvoiceDetailModel.ProductPrice)
	if err != nil {
		return nil, err
	}
	return InvoiceDetailModel, nil
}
