package request

import "github.com/gowesmart/api-gowesmart/model/web"

type CreateBikeRequest struct {
	CategoryID  uint   `json:"category_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Brand       string `json:"brand" binding:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required,url"`
	Stock       int    `json:"stock" binding:"required"`
	IsAvailable bool   `json:"is_available"`
}

type UpdateBikeRequest struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"image_url" binding:"omitempty,url"`
	Stock       int    `json:"stock"`
	IsAvailable bool   `json:"is_available"`
}

type GetBikeByIDRequest struct {
	ID uint `json:"id" binding:"required"`
}

type BikeQueryRequest struct {
	CategoryID uint   `form:"category_id" binding:"omitempty"`
	Name       string `form:"name" binding:"omitempty"`
	MinPrice   int    `form:"min_price" binding:"omitempty,gt=0"`
	MaxPrice   int    `form:"max_price" binding:"omitempty,gt=0"`
	MinYear    int    `form:"min_year" binding:"omitempty,gt=0"`
	MaxYear    int    `form:"max_year" binding:"omitempty,gt=0"`
	web.PaginationRequest
}
