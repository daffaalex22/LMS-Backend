package categories

import (
	"backend/business/categories"
	"backend/controllers"
	"backend/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoriesController struct {
	CategoriesUsecase categories.Usecase
}

func NewCategoriesController(categoryUsecase categories.Usecase) *CategoriesController {
	return &CategoriesController{
		CategoriesUsecase: categoryUsecase,
	}
}

func (categoryController CategoriesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := categoryController.CategoriesUsecase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseCategory := []response.CategoryResponse{}
	for _, value := range data {
		responseCategory = append(responseCategory, response.FromDomain(value))
	}

	return controllers.NewSuccesResponse(c, responseCategory)
}
