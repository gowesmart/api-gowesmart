package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"gorm.io/gorm"
)

type CartItemService struct{}

func NewCartItemService() *CartItemService {
	return &CartItemService{}
}

func (s CartItemService) GetByUserID(c *gin.Context, userID int) (*response.CartResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var cart entity.Cart
	if err := db.Preload("CartItem.Bike").Preload("User").Find(&cart, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	var cartItemResponse []response.CartItemResponse

	for _, val := range cart.CartItem {
		totalPrice := float64(val.Quantity) * float64(val.Bike.Price)
		cartItemResponse = append(cartItemResponse, response.CartItemResponse{
			ID:        val.ID,
			BikeID:    val.BikeID,
			CartID:    val.CartID,
			Quantity:  val.Quantity,
			Price:     totalPrice,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return &response.CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CartItems: cartItemResponse,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}

func (s CartItemService) Create(c *gin.Context, req request.CartItemCreateRequest, userID uint) (*response.CartItemResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var cart entity.Cart
	var cartItem entity.CartItem

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Select("id").Find(&cart, "user_id = ?", userID).Error
		if err != nil {
			return exceptions.NewCustomError(http.StatusBadRequest, "user not found")
		}

		cartItem = entity.CartItem{
			CartID:   cart.ID,
			BikeID:   req.BikeID,
			Quantity: req.Quantity,
		}

		if err := tx.Create(&cartItem).Error; err != nil {
			return err
		}

		return nil
	})
	utils.PanicIfError(err)

	return &response.CartItemResponse{
		ID:        cartItem.ID,
		CartID:    cartItem.CartID,
		BikeID:    cartItem.BikeID,
		Quantity:  cartItem.Quantity,
		CreatedAt: cartItem.CreatedAt,
		UpdatedAt: cartItem.UpdatedAt,
	}, nil
}
func (s CartItemService) Update(c *gin.Context, req request.CartItemUpdateRequest, userID uint) (*response.CartItemResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	cartItem := entity.CartItem{
		ID: req.ID,
		// CartID: req.CartID,
		BikeID:   req.BikeID,
		Quantity: req.Quantity,
	}

	if err := db.Save(&cartItem).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (s CartItemService) Delete(c *gin.Context, bikeID, cartID uint) error {
	db, _ := utils.GetDBAndLogger(c)

	if err := db.Where("bike_id = ?", bikeID).Where("cart_id = ?", cartID).Delete(&entity.CartItem{}).Error; err != nil {
		return err
	}

	return nil
}

// func toCartResponse(cart entity.Cart) response.CartResponse {
// 	var cartItems []response.CartItemResponse
// 	for _, item := range cart.CartItems {
// 		cartItems = append(cartItems, response.CartItemResponse{
// 			ID:       int(item.ID),
// 			BikeID:   int(item.BikeID),
// 			Quantity: item.Quantity,
// 			Price:    item.Price,
// 		})
// 	}

// 	return response.CartResponse{
// 		ID:        int(cart.ID),
// 		UserID:    int(cart.UserID),
// 		CartItems: cartItems,
// 		CreatedAt: cart.CreatedAt,
// 		UpdatedAt: cart.UpdatedAt,
// 	}
// }
