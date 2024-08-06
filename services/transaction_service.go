package services

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"gorm.io/gorm"
)

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t TransactionService) GetAll(c *gin.Context) ([]response.TransactionResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var transactions []entity.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	var results []response.TransactionResponse
	for _, transaction := range transactions {
		result := toResponse(transaction)
		results = append(results, result)
	}

	return results, nil
}

func (t TransactionService) GetById(c *gin.Context, transactionId int) (response.TransactionResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var transaction entity.Transaction
	if err := db.Preload("Order").Where("id = ?", transactionId).First(&transaction).Error; err != nil {
		return response.TransactionResponse{}, err
	}

	result := toResponse(transaction)

	return result, nil
}

func (t TransactionService) Create(c *gin.Context, payloads []request.TransactionCreate, userId int) error {
	db, _ := utils.GetDBAndLogger(c)
	var wg sync.WaitGroup
	channels := make(chan error, len(payloads))

	err := db.Transaction(func(tx *gorm.DB) error {
		transaction := toTransactionEntity(userId, payloads)

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		wg.Add(len(payloads))
		for _, payload := range payloads {
			go createOrder(&wg, channels, tx, userId, transaction.ID, payload)
		}

		wg.Wait()
		close(channels)

		for channel := range channels {
			if channel != nil {
				return channel
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionService) Update(c *gin.Context, payloads []request.TransactionUpdate, transactionId int) error {
	db, _ := utils.GetDBAndLogger(c)
	var wg sync.WaitGroup
	channels := make(chan error, len(payloads))

	err := db.Transaction(func(tx *gorm.DB) error {
		var transaction entity.Transaction
		if err := tx.Where("id = ?", transactionId).First(&transaction).Error; err != nil {
			return err
		}

		if transaction.Status != "pending" {
			return errors.New("invalid transaction")
		}

		wg.Add(len(payloads))
		for _, payload := range payloads {
			go updateorder(&wg, channels, tx, payload, &transaction)
		}

		wg.Wait()
		close(channels)

		for channel := range channels {
			if channel != nil {
				return channel
			}
		}

		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionService) Delete(c *gin.Context, transactionId int) error {
	db, _ := utils.GetDBAndLogger(c)

	var transaction entity.Transaction
	if err := db.Where("id = ?", transactionId).Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (t TransactionService) Pay(c *gin.Context, transactionId int) error {
	db, _ := utils.GetDBAndLogger(c)

	var transaction entity.Transaction
	if err := db.Where("id = ?", transactionId).First(&transaction).Error; err != nil {
		return err
	}

	if transaction.Status != "pending" {
		return errors.New("invalid transaction")
	}

	transaction.Status = "paid"
	if err := db.Save(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (t TransactionService) GetTransactionByUserID(c *gin.Context, userId uint) (*response.UserTransactionResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var user entity.User

	err := db.Preload("Transaction").Take(&user, userId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "user not found")
		}
		return nil, err
	}

	var userTransactions []response.TransactionResponse
	for _, transaction := range user.Transaction {
		result := toResponse(transaction)
		userTransactions = append(userTransactions, result)
	}

	return &response.UserTransactionResponse{
		ID:          user.ID,
		Username:    user.Username,
		Transaction: userTransactions,
	}, nil
}

// helpers
func toTransactionEntity(userId int, payloads []request.TransactionCreate) entity.Transaction {
	transaction := entity.Transaction{
		UserID: userId,
		Status: "pending",
	}

	for _, payload := range payloads {
		transaction.TotalPrice += payload.TotalPrice
	}

	return transaction
}

func toOrderEntity(userId int, transactionId int, payload request.TransactionCreate) entity.Order {
	order := entity.Order{
		BikeID:        payload.BikeID,
		Quantity:      payload.Quantity,
		TotalPrice:    payload.TotalPrice,
		UserID:        userId,
		TransactionID: transactionId,
	}

	return order
}

func toResponse(payload entity.Transaction) response.TransactionResponse {
	var orders []response.OrderResponse

	for _, order := range payload.Order {
		temp := response.OrderResponse{
			ID:         order.ID,
			BikeID:     order.BikeID,
			Quantity:   order.Quantity,
			TotalPrice: order.TotalPrice,
		}

		orders = append(orders, temp)
	}

	return response.TransactionResponse{
		ID:         payload.ID,
		TotalPrice: payload.TotalPrice,
		UserID:     payload.UserID,
		Status:     payload.Status,
		Orders:     orders,
	}
}

// concurrent
func createOrder(wg *sync.WaitGroup, channels chan error, tx *gorm.DB, userId, transactionId int, payload request.TransactionCreate) {
	defer wg.Done()

	order := toOrderEntity(userId, transactionId, payload)
	if err := tx.Create(&order).Error; err != nil {
		channels <- err
		return
	}

	channels <- nil
}

func updateorder(wg *sync.WaitGroup, channels chan error, tx *gorm.DB, payload request.TransactionUpdate, transaction *entity.Transaction) {
	defer wg.Done()

	var order entity.Order
	if err := tx.Where("id = ?", payload.ID).First(&order).Error; err != nil {
		channels <- err
		return
	}

	transaction.TotalPrice -= order.TotalPrice
	transaction.TotalPrice += payload.TotalPrice

	order.BikeID = payload.BikeID
	order.Quantity = payload.Quantity
	order.TotalPrice = payload.TotalPrice

	if err := tx.Save(&order).Error; err != nil {
		channels <- err
		return
	}

	channels <- nil
}
