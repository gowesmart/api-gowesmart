package services

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t TransactionService) GetAll(c *gin.Context, paginationReq *web.PaginationRequest) ([]response.TransactionResponse, *web.Metadata, error) {
	db, logger := utils.GetDBAndLogger(c)

	var transactions []entity.Transaction

	var totalData int64
	query := db.Model(&entity.Transaction{})
	if err := query.Count(&totalData).Error; err != nil {
		logger.Error("failed to count transactions", zap.Error(err))
		return nil, nil, err
	}
	paginationReq.TotalData = totalData

	offset := paginationReq.GetOffset()
	limit := paginationReq.GetLimit()
	if err := query.Offset(offset).Limit(limit).Preload("Order").Find(&transactions).Error; err != nil {
		logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, nil, err
	}

	paginationReq.TotalPages = int((totalData + int64(limit) - 1) / int64(limit))

	var results []response.TransactionResponse
	for _, transaction := range transactions {
		results = append(results, toResponse(transaction))
	}

	metadata := &web.Metadata{
		Page:       &paginationReq.Page,
		Limit:      &paginationReq.Limit,
		TotalPages: &paginationReq.TotalPages,
		TotalData:  &paginationReq.TotalData,
	}

	logger.Info("success fetching all transactions", zap.Int("total_data", int(totalData)), zap.Int("total_pages", paginationReq.TotalPages))

	return results, metadata, nil
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

func (t TransactionService) Create(c *gin.Context, payloads []request.TransactionCreate, userID int) (response.CreateTransactionResponse, error) {
	db, _ := utils.GetDBAndLogger(c)
	var wg sync.WaitGroup
	channels := make(chan error, len(payloads))
	var response response.CreateTransactionResponse

	err := db.Transaction(func(tx *gorm.DB) error {
		transaction := toTransactionEntity(userID, payloads)

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		wg.Add(len(payloads))
		for _, payload := range payloads {
			go createOrder(&wg, channels, tx, userID, transaction.ID, payload)
		}

		wg.Wait()
		close(channels)

		for channel := range channels {
			if channel != nil {
				return channel
			}
		}

		response.TransactionID = transaction.ID
		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}

func (t TransactionService) Update(c *gin.Context, payloads []request.TransactionUpdate, transactionID, userID int) error {
	db, _ := utils.GetDBAndLogger(c)
	var wg sync.WaitGroup
	channels := make(chan error, len(payloads))

	err := db.Transaction(func(tx *gorm.DB) error {
		var transaction entity.Transaction
		if err := tx.Where("user_id = ?", userID).Where("id = ?", transactionID).First(&transaction).Error; err != nil {
			return err
		}

		if transaction.Status != "pending" {
			return exceptions.NewCustomError(http.StatusBadRequest, "Invalid transaction")
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

func (t TransactionService) Delete(c *gin.Context, transactionID, userID int) error {
	db, _ := utils.GetDBAndLogger(c)

	var transaction entity.Transaction
	if err := db.Where("user_id = ?", userID).Where("id = ?", transactionID).Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (t TransactionService) Pay(c *gin.Context, transactionID, userID int) error {
	db, _ := utils.GetDBAndLogger(c)

	var transaction entity.Transaction
	if err := db.Where("user_id = ?", userID).Where("id = ?", transactionID).First(&transaction).Error; err != nil {
		return err
	}

	if transaction.Status != "pending" {
		return exceptions.NewCustomError(http.StatusBadRequest, "Invalid transaction")
	}

	transaction.Status = "paid"
	if err := db.Save(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (t TransactionService) GetTransactionByUserID(c *gin.Context, userID uint) (*response.UserTransactionResponse, error) {
	db, _ := utils.GetDBAndLogger(c)

	var user entity.User

	err := db.Preload("Transaction").Take(&user, userID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "User not found")
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
