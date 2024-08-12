package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BikeService struct{}

func NewBikeService() *BikeService {
	return &BikeService{}
}

func (service *BikeService) CreateBike(c *gin.Context, bikeReq *request.CreateBikeRequest) (*response.BikeResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.BikeResponse

	bike := entity.Bike{
		CategoryID:  bikeReq.CategoryID,
		Name:        bikeReq.Name,
		Brand:       bikeReq.Brand,
		Description: bikeReq.Description,
		Year:        bikeReq.Year,
		Price:       bikeReq.Price,
		ImageUrl:    bikeReq.ImageUrl,
		Stock:       bikeReq.Stock,
		IsAvailable: bikeReq.IsAvailable,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bike).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Bike{}).
			Select("id, category_id, name, brand, description, year, price, image_url, stock, is_available, created_at, updated_at").
			Take(&res, bike.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success creating bike", zap.Uint("bikeID", bike.ID))

	return &res, nil
}

func (service *BikeService) UpdateBike(c *gin.Context, id uint, bikeReq *request.UpdateBikeRequest) (*response.BikeResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.BikeResponse
	var bike entity.Bike

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&bike, id).Error; err != nil {
			return err
		}

		if bikeReq.CategoryID != 0 {
			bike.CategoryID = bikeReq.CategoryID
		}
		if bikeReq.Name != "" {
			bike.Name = bikeReq.Name
		}
		if bikeReq.Brand != "" {
			bike.Brand = bikeReq.Brand
		}
		if bikeReq.Description != "" {
			bike.Description = bikeReq.Description
		}
		if bikeReq.Year != 0 {
			bike.Year = bikeReq.Year
		}
		if bikeReq.Price != 0 {
			bike.Price = bikeReq.Price
		}
		if bikeReq.ImageUrl != "" {
			bike.ImageUrl = bikeReq.ImageUrl
		}
		if bikeReq.Stock != 0 {
			bike.Stock = bikeReq.Stock
		}
		bike.IsAvailable = bikeReq.IsAvailable

		if err := tx.Save(&bike).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Bike{}).
			Select("id, category_id, name, brand, description, year, price, image_url, stock, is_available, created_at, updated_at").
			Take(&res, bike.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating bike", zap.Uint("bikeID", bike.ID))

	return &res, nil
}

func (service *BikeService) DeleteBike(c *gin.Context, id uint) error {
	db, logger := utils.GetDBAndLogger(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.Bike{}, id).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("success deleting bike", zap.Uint("bikeID", id))

	return nil
}

func (service *BikeService) GetAllBikes(c *gin.Context, bikeQueryReq *request.BikeQueryRequest) ([]response.BikeResponse, *web.Metadata, error) {
	db, logger := utils.GetDBAndLogger(c)

	var bikes []response.BikeResponse

	query := db.Model(&entity.Bike{}).
		Where("stock > 0").
		Select("id, category_id, name, brand, description, year, price, image_url, stock, is_available, created_at, updated_at")

	if bikeQueryReq.CategoryID != 0 {
		query = query.Where("category_id = ?", bikeQueryReq.CategoryID)
	}
	if bikeQueryReq.Name != "" {
		query = query.Where("to_tsvector('english', name) @@ plainto_tsquery('english', ?)", bikeQueryReq.Name)
	}
	if bikeQueryReq.MinPrice > 0 {
		query = query.Where("price >= ?", bikeQueryReq.MinPrice)
	}
	if bikeQueryReq.MaxPrice > 0 {
		query = query.Where("price <= ?", bikeQueryReq.MaxPrice)
	}
	if bikeQueryReq.MinYear > 0 {
		query = query.Where("year >= ?", bikeQueryReq.MinYear)
	}
	if bikeQueryReq.MaxYear > 0 {
		query = query.Where("year <= ?", bikeQueryReq.MaxYear)
	}

	var totalData int64
	if err := query.Count(&totalData).Error; err != nil {
		logger.Error("failed to count bikes", zap.Error(err))
		return nil, nil, err
	}
	bikeQueryReq.TotalData = totalData

	offset := bikeQueryReq.GetOffset()
	limit := bikeQueryReq.GetLimit()
	if err := query.Offset(offset).Limit(limit).Find(&bikes).Error; err != nil {
		logger.Error("failed to fetch bikes", zap.Error(err))
		return nil, nil, err
	}

	bikeQueryReq.TotalPages = int((totalData + int64(limit) - 1) / int64(limit))

	metadata := &web.Metadata{
		Page:       &bikeQueryReq.Page,
		Limit:      &bikeQueryReq.Limit,
		TotalPages: &bikeQueryReq.TotalPages,
		TotalData:  &bikeQueryReq.TotalData,
	}

	logger.Info("success fetching all bikes", zap.Int("total_data", int(totalData)), zap.Int("total_pages", bikeQueryReq.TotalPages))

	return bikes, metadata, nil
}

func (service *BikeService) GetBikeByID(c *gin.Context, id uint) (*response.BikeResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.BikeResponse

	if err := db.Model(&entity.Bike{}).
		Select("id, category_id, name, brand, description, year, price, image_url, stock, is_available, created_at, updated_at").
		Take(&res, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("bike not found", zap.Uint("bikeID", id))
			return nil, err
		}

		logger.Error("failed to fetch bike", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching bike", zap.Uint("bikeID", id))

	return &res, nil
}
