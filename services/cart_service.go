package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
)

type CartService struct{}

func NewCartService() *CartService {
	return &CartService{}
}

func (s CartService) GetAll(c *gin.Context) ([]response.CartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var carts []entity.Cart
	if err := db.Preload("CartItems").Find(&carts).Error; err != nil {
		return nil, err
	}

	var results []response.CartResponse
	for _, cart := range carts {
		result := toCartResponse(cart)
		results = append(results, result)
	}

	return results, nil
}

func (s CartService) GetById(c *gin.Context, cartId int) (response.CartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var cart entity.Cart
	if err := db.Preload("CartItems").Where("id = ?", cartId).First(&cart).Error; err != nil {
		return response.CartResponse{}, err
	}

	result := toCartResponse(cart)
	return result, nil
}

func (s CartService) Create(c *gin.Context, req request.CartCreateRequest) (response.CartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	cart := entity.Cart{
		UserID: req.UserID,
	}

	if err := db.Create(&cart).Error; err != nil {
		return response.CartResponse{}, err
	}

	return toCartResponse(cart), nil
}

func (s CartService) Update(c *gin.Context, cartId int, req request.CartUpdateRequest) (response.CartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var cart entity.Cart
	if err := db.Where("id = ?", cartId).First(&cart).Error; err != nil {
		return response.CartResponse{}, err
	}

	cart.UserID = req.UserID

	if err := db.Save(&cart).Error; err != nil {
		return response.CartResponse{}, err
	}

	return toCartResponse(cart), nil
}

func (s CartService) Delete(c *gin.Context, cartId int) error {
	db, _ := utils.GetDBAndLogger(c)

	if err := db.Where("id = ?", cartId).Delete(&entity.Cart{}).Error; err != nil {
		return err
	}

	return nil
}

func toCartResponse(cart entity.Cart) response.CartResponse {
	var cartItems []response.CartItemResponse
	for _, item := range cart.CartItems {
		cartItems = append(cartItems, response.CartItemResponse{
			ID:       item.ID,
			BikeID:   item.BikeID,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}

	return response.CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CartItems: cartItems,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}
