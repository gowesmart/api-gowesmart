package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/entity"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	"github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (service *CategoryService) CreateCategory(c *gin.Context, categoryReq *request.CreateCategoryRequest) (*response.CategoryResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.CategoryResponse

	category := entity.Category{
		Name: categoryReq.Name,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&category).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Category{}).
			Select("id, name").
			Take(&res, category.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success creating category", zap.Uint("categoryID", category.ID))

	return &res, nil
}

func (service *CategoryService) UpdateCategory(c *gin.Context, id uint, categoryReq *request.UpdateCategoryRequest) (*response.CategoryResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.CategoryResponse
	var category entity.Category

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&category, id).Error; err != nil {
			return err
		}

		category.Name = categoryReq.Name

		if err := tx.Save(&category).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Category{}).
			Select("id, name").
			Take(&res, category.ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating category", zap.Uint("categoryID", category.ID))

	return &res, nil
}

func (service *CategoryService) DeleteCategory(c *gin.Context, id uint) error {
	db, logger := utils.GetDBAndLogger(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entity.Category{}, id).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("success deleting category", zap.Uint("categoryID", id))

	return nil
}

func (service *CategoryService) GetAllCategories(c *gin.Context) ([]response.CategoryResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var categories []response.CategoryResponse

	if err := db.Model(&entity.Category{}).
		Select("id, name").
		Find(&categories).Error; err != nil {
		logger.Error("failed to fetch categories", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching all categories")

	return categories, nil
}

func (service *CategoryService) GetCategoryByID(c *gin.Context, id uint) (*response.CategoryResponse, error) {
	db, logger := utils.GetDBAndLogger(c)

	var res response.CategoryResponse

	if err := db.Model(&entity.Category{}).
		Select("id, name").
		Take(&res, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warn("category not found", zap.Uint("categoryID", id))
			return nil, err
		}

		logger.Error("failed to fetch category", zap.Error(err))
		return nil, err
	}

	logger.Info("success fetching category", zap.Uint("categoryID", id))

	return &res, nil
}
