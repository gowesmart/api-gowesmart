package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartItemService struct{}

func NewCartItemService() *CartItemService {
	return &CartItemService{}
}

func (s CartItemService) GetByUserID(c *gin.Context, userID uint) (*response.GetUserCartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var cart entity.Cart
	if err := db.Preload("CartItem.Bike").Preload("User").Find(&cart, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	var cartItemResponse []response.GetUserCartItemResponse

	for _, val := range cart.CartItem {
		totalPrice := float64(val.Quantity) * float64(val.Bike.Price)
		cartItemResponse = append(cartItemResponse, response.GetUserCartItemResponse{
			ID: val.ID,
			Bike: response.GetUserCartItemBikeResponse{
				ID:          val.Bike.ID,
				Brand:       val.Bike.Brand,
				Name:        val.Bike.Name,
				Price:       val.Bike.Price,
				ImageUrl:    val.Bike.ImageUrl,
				Stock:       val.Bike.Stock,
				Description: val.Bike.Description,
			},
			CartID:    val.CartID,
			Quantity:  val.Quantity,
			Price:     totalPrice,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return &response.GetUserCartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CartItems: cartItemResponse,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}

func (s CartItemService) Create(c *gin.Context, req request.CartItemCreateRequest, userID uint) (*response.CartItemResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var cart entity.Cart
	var cartItem entity.CartItem

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Select("id").Where("user_id = ?", userID).First(&cart).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusBadRequest, "User not found")
		}

		if err := tx.Where("bike_id = ? AND cart_id = ?", req.BikeID, cart.ID).First(&cartItem).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				cartItem = entity.CartItem{
					CartID:   cart.ID,
					BikeID:   req.BikeID,
					Quantity: req.Quantity,
				}
				if err := tx.Create(&cartItem).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			cartItem.Quantity += req.Quantity
			if err := tx.Save(&cartItem).Error; err != nil {
				return err
			}
		}

		logger.Info("success creating or updating cart item", zap.Uint("cartItemID", cartItem.ID))

		return nil
	})
	utils.PanicIfError(err)

	return s.toCartItemResponse(cartItem), nil
}

func (s CartItemService) Update(c *gin.Context, req request.CartItemUpdateRequest, userID uint) (*response.CartItemResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var cart entity.Cart
	var cartItem entity.CartItem

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Select("id").Where("user_id = ?", userID).First(&cart).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusBadRequest, "User not found")
		}

		if err := tx.Where("bike_id = ? AND cart_id = ?", req.BikeID, cart.ID).First(&cartItem).Error; err != nil {
			return exceptions.NewCustomError(http.StatusNotFound, "Cart item not found")
		}

		cartItem.Quantity = req.Quantity
		if err := tx.Save(&cartItem).Error; err != nil {
			return err
		}

		logger.Info("success updating cart item", zap.Uint("cartItemID", cartItem.ID))

		return nil
	})
	utils.PanicIfError(err)

	return s.toCartItemResponse(cartItem), nil
}
func (s CartItemService) Delete(c *gin.Context, bikeID, userID uint) error {
	db, logger := utils.GetDBAndLogger(c)

	var cart entity.Cart

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Select("id").Where("user_id = ?", userID).First(&cart).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusBadRequest, "User not found")
		}

		result := tx.Where("bike_id = ? AND cart_id = ?", bikeID, cart.ID).Delete(&entity.CartItem{})
		if result.Error != nil {
			return err
		}

		if result.RowsAffected == 0 {
			return exceptions.NewCustomError(http.StatusNotFound, "Cart item not found")
		}

		logger.Info("success deleting cart item", zap.Uint("cartID", cart.ID), zap.Uint("bikeID", bikeID))

		return nil
	})
	utils.PanicIfError(err)

	return nil
}

func (s CartItemService) toCartItemResponse(cartItem entity.CartItem) *response.CartItemResponse {
	return &response.CartItemResponse{
		ID:        cartItem.ID,
		CartID:    cartItem.CartID,
		BikeID:    cartItem.BikeID,
		Quantity:  cartItem.Quantity,
		CreatedAt: cartItem.CreatedAt,
		UpdatedAt: cartItem.UpdatedAt,
	}
}
