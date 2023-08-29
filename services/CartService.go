package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
	"fmt"
)

type CartServive struct {
}

var CartDao dao.CartDAO

func (CartServive *CartServive) AddToCart(userId, productId int) error {
	query := fmt.Sprintf("WHERE user_id = %d AND product_id = %d", userId, productId)
	result, err := CartDao.FindByCondition(query)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		var cartModel models.CartModel
		cartModel.UserId = userId
		cartModel.ProductId = productId
		cartModel.Quantity = 1
		CartDao.Create(cartModel)
		return nil
	} else {
		var cartModel models.CartModel
		cartModel.UserId = userId
		cartModel.ProductId = productId
		cartModel.Quantity = result[0].Quantity + 1
		CartDao.Update(cartModel)
		return nil
	}
}

func (CartServive *CartServive) UpdateInCart(userId, productId int, types string) error {
	if types == "increase" {
		CartServive.AddToCart(userId, productId)
		return nil
	} else if types == "decrease" {
		query := fmt.Sprintf("WHERE user_id = %d AND product_id = %d", userId, productId)
		result, err := CartDao.FindByCondition(query)
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("Không tìm thấy trong giỏ hàng")
		} else {
			cartModel := result[0]
			if cartModel.Quantity == 1 {
				CartDao.Delete(userId, productId)
				return nil
			} else {
				cartModel.Quantity = cartModel.Quantity - 1
				CartDao.Update(cartModel)
				return nil
			}
		}
	}
	return errors.New("Loại không tồn tại")
}

func (CartServive *CartServive) DeleteInCart(userId, productId int) error {
	query := fmt.Sprintf("WHERE user_id = %d AND product_id = %d", userId, productId)
	result, err := CartDao.FindByCondition(query)
	if err != nil {
		return errors.New("Không tìm thấy trong giỏ hàng")
	}
	if len(result) == 1 {
		CartDao.Delete(userId, productId)
		return nil
	} else {
		return errors.New("Không có sản phẩm để xóa")
	}
}
