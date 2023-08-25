package dao

import (
	"computer_shop/config"
	"computer_shop/models"
)

type CategoryDAO struct {
}

func (CategoryDAO *CategoryDAO) FindAll() ([]models.CategoryModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Categories";
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var categories []models.CategoryModel
	for rows.Next() {
		var category models.CategoryModel
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}


func (CategoryDAO *CategoryDAO) FindById(id int) (models.CategoryModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Categories WHERE category_id = ?"
	var category models.CategoryModel
	if err := db.QueryRow(query, id).Scan(&category.CategoryId, &category.CategoryName, &category.CreatedAt, &category.UpdatedAt); err != nil {
		return category, err
	}
	return category, nil
}

//return -1 == error
func (CategoryDAO *CategoryDAO) Create(category models.CategoryModel) (int, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "INSERT INTO Categories(category_name) VALUES (?)"
	result, err := db.Exec(query, category.CategoryName)
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (CategoryDAO *CategoryDAO) Update(category models.CategoryModel) error {
	db := config.GetConnection()
	defer db.Close()
	query := "UPDATE Categories SET category_name = ? WHERE category_id = ?"
	_, err := db.Exec(query, category.CategoryName, category.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func (CategoryDAO *CategoryDAO) Delete(id int) error {
	db := config.GetConnection()
	defer db.Close()
	query := "DELETE FROM Categories WHERE category_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (CategoryDAO *CategoryDAO) FindByCondition(condition string) ([]models.CategoryModel, error) {
	db := config.GetConnection()
	defer db.Close()
	query := "SELECT * FROM Categories " + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var categories []models.CategoryModel
	for rows.Next() {
		var category models.CategoryModel
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

