package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/exceptions"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/model/web/request"
	_ "github.com/gowesmart/api-gowesmart/model/web/response"
	"github.com/gowesmart/api-gowesmart/services"
	"github.com/gowesmart/api-gowesmart/utils"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		*categoryService,
	}
}

// CreateCategory godoc
// @Summary Create a category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param category body request.CreateCategoryRequest true "Category body"
// @Success 201 {object} web.WebSuccess[response.CategoryResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/categories [post]
func (controller *CategoryController) CreateCategory(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	var categoryReq request.CreateCategoryRequest
	err := c.ShouldBindJSON(&categoryReq)
	utils.PanicIfError(err)

	res, err := controller.categoryService.CreateCategory(c, &categoryReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusCreated, res, nil)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description	Update an existing category
// @Tags Categories
// @Accept json
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true "Category ID"
// @Param category body request.UpdateCategoryRequest	true "Category body"
// @Success 200	{object} web.WebSuccess[response.CategoryResponse]
// @Failure 400	{object} web.WebBadRequestError
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/categories/{id} [patch]
func (controller *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	utils.UserRoleMustAdmin(c)

	var categoryReq request.UpdateCategoryRequest

	err = c.ShouldBindJSON(&categoryReq)
	utils.PanicIfError(err)

	res, err := controller.categoryService.UpdateCategory(c, uint(id), &categoryReq)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description	Delete a category by ID
// @Tags Categories
// @Produce json
// @Param Authorization	header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path uint true	"Category ID"
// @Success 204	{object} web.WebSuccess[response.CategoryResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/categories/{id} [delete]
func (controller *CategoryController) DeleteCategory(c *gin.Context) {
	utils.UserRoleMustAdmin(c)

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.PanicIfError(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	err = controller.categoryService.DeleteCategory(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, "Category deleted", nil)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description	Get all categories
// @Tags Categories
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Success 200	{object} web.WebSuccess[[]response.CategoryResponse]
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/categories [get]
func (controller *CategoryController) GetAllCategories(c *gin.Context) {
	var pagination web.PaginationRequest

	err := c.ShouldBindQuery(&pagination)
	utils.PanicIfError(err)

	res, metadata, err := controller.categoryService.GetAllCategories(c, &pagination)
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, metadata)
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description	Get a category by ID
// @Tags Categories
// @Produce json
// @Param id path uint true	"Category ID"
// @Success 200	{object} web.WebSuccess[response.CategoryResponse]
// @Failure 404	{object} web.WebNotFoundError
// @Failure 500	{object} web.WebInternalServerError
// @Router /api/categories/{id} [get]
func (controller *CategoryController) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	res, err := controller.categoryService.GetCategoryByID(c, uint(id))
	utils.PanicIfError(err)

	utils.ToResponseJSON(c, http.StatusOK, res, nil)
}
