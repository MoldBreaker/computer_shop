package models

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
	ProductId    int
	InvoiceId    int
	Quantity     int
	ProductPrice int
}

type InvoiceDetailResponseModel struct {
	Product      ProductResponse
	Quantity     int
	ProductPrice int
}

func (InvoiceDetailModel *InvoiceDetailModel) GetTableName() string {
	return `Invoice_Detail`
}
