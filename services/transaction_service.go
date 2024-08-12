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

func (t TransactionService) GetAll(c *gin.Context, paginationReq *web.PaginationRequest) ([]response.GetAllTransactionResponse, *web.Metadata, error) {
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
	if err := query.Offset(offset).
		Limit(limit).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id, username") }).
		Preload("Order.Bike").
		Find(&transactions).Error; err != nil {
		logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, nil, err
	}

	paginationReq.TotalPages = int((totalData + int64(limit) - 1) / int64(limit))

	var results []response.GetAllTransactionResponse
	for _, transaction := range transactions {
		results = append(results, toGetAllResponse(transaction))
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

		var user entity.User
		if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
			return err
		}

		paymentPayload := utils.PaymentPayload{
			OrderId: transaction.ID,
			Amount:  transaction.TotalPrice,
			FName:   user.Username,
			Email:   user.Email,
		}

		paymentLink, err := utils.CreatePayment(paymentPayload)
		if err != nil {
			return err
		}

		transaction.PaymentLink = paymentLink
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

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

	err := db.Transaction(func(tx *gorm.DB) error {
		var transaction entity.Transaction
		if err := tx.Preload("Order").Where("user_id = ?", userID).Where("id = ?", transactionID).First(&transaction).Error; err != nil {
			return err
		}

		if transaction.Status != "pending" {
			return exceptions.NewCustomError(http.StatusBadRequest, "Invalid transaction")
		}

		transaction.Status = "paid"
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		for _, order := range transaction.Order {
			var bike entity.Bike
			if err := tx.Where("id = ?", order.BikeID).First(&bike).Error; err != nil {
				return err
			}

			bike.Stock -= order.Quantity
			if err := tx.Save(&bike).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionService) GetTransactionByUserID(c *gin.Context, paginationReq *web.PaginationRequest, userID uint) ([]response.GetAllTransactionResponse, *web.Metadata, error) {
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
	if err := query.Offset(offset).
		Limit(limit).
		Where("user_id = ?", userID).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id, username") }).
		Preload("Order.Bike").
		Find(&transactions).Error; err != nil {
		logger.Error("failed to fetch transactions", zap.Error(err))
		return nil, nil, err
	}

	paginationReq.TotalPages = int((totalData + int64(limit) - 1) / int64(limit))

	var results []response.GetAllTransactionResponse
	for _, transaction := range transactions {
		results = append(results, toGetAllResponse(transaction))
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
		ID:          payload.ID,
		TotalPrice:  payload.TotalPrice,
		UserID:      payload.UserID,
		Status:      payload.Status,
		PaymentLink: payload.PaymentLink,
		Orders:      orders,
		CreatedAt:   payload.CreatedAt.Format("02-01-2006"),
		UpdatedAt:   payload.UpdatedAt.Format("02-01-2006"),
	}
}

func toGetAllResponse(payload entity.Transaction) response.GetAllTransactionResponse {
	var orders []response.GetAllOrderResponse

	for _, order := range payload.Order {
		temp := response.GetAllOrderResponse{
			ID: order.ID,
			Bike: response.GetAllOrderBikeResponse{
				ID:       order.Bike.ID,
				Name:     order.Bike.Name,
				ImageUrl: order.Bike.ImageUrl,
			},
			Quantity:   order.Quantity,
			TotalPrice: order.TotalPrice,
		}

		orders = append(orders, temp)
	}

	return response.GetAllTransactionResponse{
		ID:         payload.ID,
		TotalPrice: payload.TotalPrice,
		User: response.GetAllTransactionUserResponse{
			ID:       payload.User.ID,
			Username: payload.User.Username,
		},
		Status:    payload.Status,
		Orders:    orders,
		CreatedAt: payload.CreatedAt,
		UpdatedAt: payload.UpdatedAt,
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
