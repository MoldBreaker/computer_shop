package dao

import (
	"computer_shop/models"
)

type AbstractDAO[T models.Models] struct {
	Data T
}

//func (AbstractDAO *AbstractDAO[T]) FindAll(query string) ([]T, error) {
//	db := config.GetConnection()
//	defer db.Close()
//	rows, err := db.Query(query)
//	defer rows.Close()
//	if err != nil {
//		return nil, err
//	}
//	var arr []T
//	for rows.Next() {
//		var result T
//		switch result.GetTableName() {
//		case "Categories":
//			result, err = result.ScanToCategoryModel(rows)
//
//		}
//	}
//	return arr, nil
//}
