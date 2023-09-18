package dao

import (
	"computer_shop/config"
	"computer_shop/models"
)

type CartDAO struct{}

/*
CREATE TABLE `carts` (

	`user_id` int NOT NULL,
	`product_id` int NOT NULL,
	`quantity` int DEFAULT NULL,
	PRIMARY KEY (`user_id`,`product_id`),
	KEY `cart_product_1212_idx` (`product_id`),
	KEY `cart_users+1274_idx` (`user_id`),
	CONSTRAINT `cart_product_1212` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`),
	CONSTRAINT `cart_users+1274` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)

)
*/

func (c CartDAO) Create(cart models.CartModel) (int, error) {
	db := config.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO carts (user_id, product_id, quantity) VALUES (?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(cart.UserId, cart.ProductId, cart.Quantity)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (c CartDAO) FindAll() ([]models.CartModel, error) {
	db := config.GetConnection()
	defer db.Close()
	var carts []models.CartModel
	rows, err := db.Query("SELECT * FROM carts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart models.CartModel
		err := rows.Scan(&cart.UserId, &cart.ProductId, &cart.Quantity)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (c CartDAO) FindById(userId int, productId int) (models.CartModel, error) {
	db := config.GetConnection()
	defer db.Close()
	var cart models.CartModel
	err := db.QueryRow("SELECT * FROM carts WHERE user_id =? AND product_id = ?", userId, productId).Scan(&cart.UserId, &cart.ProductId, &cart.Quantity)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (c CartDAO) Update(cart models.CartModel) error {
	db := config.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE carts SET quantity =? WHERE user_id =? AND product_id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cart.Quantity, cart.UserId, cart.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (c CartDAO) Delete(userId int, productId int) error {
	db := config.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM carts WHERE user_id =? AND product_id =?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c CartDAO) FindByCondition(condition string) ([]models.CartModel, error) {
	db := config.GetConnection()
	defer db.Close()
	var carts []models.CartModel
	rows, err := db.Query("SELECT * FROM carts " + condition)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart models.CartModel
		err := rows.Scan(&cart.UserId, &cart.ProductId, &cart.Quantity)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
