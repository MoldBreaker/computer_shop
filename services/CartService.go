package services

import (
	"computer_shop/dao"
	"computer_shop/models"
	"errors"
	"fmt"
)

type CartService struct {
}

var (
	CartDao dao.CartDAO
)

func (CartService *CartService) AddToCart(userId, productId int) error {
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

func (CartService *CartService) UpdateInCart(userId, productId int, types string) error {
	if types == "increase" {
		CartService.AddToCart(userId, productId)
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

func (CartService *CartService) DeleteInCart(userId, productId int) error {
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

func (CartService *CartService) GetCartByUserId(userId int) ([]models.CartResponseModel, error) {
	condition := fmt.Sprintf("WHERE user_id = %d", userId)
	cart, err := CartDao.FindByCondition(condition)
	if err != nil {
		return nil, errors.New("Error when getting Cart")
	}
	if len(cart) == 0 {
		return nil, errors.New("You haven't added anything to your cart yet")
	}
	var cartResponse []models.CartResponseModel
	for _, c := range cart {
		product, err := ProductDAO.FindById(c.ProductId)
		if err != nil {
			return nil, errors.New("Error when getting Product")
		}
		var productRes models.ProductResponse
		var ProductImageService ProductImageService
		urls, err := ProductImageService.GetImagesByProductId(product.ProductId)
		productRes = productRes.Parse(product, urls)

		item := models.CartResponseModel{
			Product:  productRes,
			Price:    c.Quantity * productRes.Price,
			Quantity: c.Quantity,
		}
		cartResponse = append(cartResponse, item)
	}
	return cartResponse, nil
}

func (CartService *CartService) DeleteAllCart(userId int) error {
	items, err := CartService.GetCartByUserId(userId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error when getting Cart"))
	}
	for i := 0; i < len(items); i++ {
		if err := CartDao.Delete(userId, items[i].Product.ProductId); err != nil {
			return errors.New(fmt.Sprintf("Error when deleting product"))
		}
	}
	return nil
}
