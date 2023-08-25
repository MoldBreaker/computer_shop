package models

import "database/sql"

type Models interface {
	*CategoryModel | *InvoiceModel | *InvoiceDetailModel | *NotificationModel | *ProductImageModel | *ProductModel | *RoleModel | *UserModel
	GetTableName() string
	ScanToCategoryModel(rows *sql.Rows) (*CategoryModel, error)
	ScanToInvoiceModel(rows *sql.Rows) (*InvoiceModel, error)
	ScanToInvoiceDetailModel(rows *sql.Rows) (*InvoiceDetailModel, error)
	ScanToNotificationModel(rows *sql.Rows) (*NotificationModel, error)
	ScanToProductImageModel(rows *sql.Rows) (*ProductImageModel, error)
	ScanToProductModel(rows *sql.Rows) (*ProductModel, error)
	ScanToRoleModel(rows *sql.Rows) (*RoleModel, error)
	ScanToUserModel(rows *sql.Rows) (*UserModel, error)
}
