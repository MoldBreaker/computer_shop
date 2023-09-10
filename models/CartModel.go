package models

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
type CartModel struct {
	UserId    int
	ProductId int
	Quantity  int
}

type CartResponseModel struct {
	Product  ProductResponse
	Price    int
	Quantity int
}

func (CartModel *CartModel) GetTablename() string {
	return "carts"
}
