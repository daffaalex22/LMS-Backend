package difficulties

import (
	"backend/business/difficulties"
	"backend/controllers"
	"backend/controllers/difficulties/response"
	"backend/helper/err"
	"fmt"

	"github.com/labstack/echo/v4"
)

type DifficultiesController struct {
	DifficultiesUsecase difficulties.Usecase
}

func NewDifficultiesController(difficultyUsecase difficulties.Usecase) *DifficultiesController {
	return &DifficultiesController{
		DifficultiesUsecase: difficultyUsecase,
	}
}

func (difficultyController DifficultiesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, getErr := difficultyController.DifficultiesUsecase.GetAll(ctx)

	if getErr != nil {
		errCode := err.ErrorDifficultyCheck(getErr)
		fmt.Println(errCode)
		return controllers.ErrorResponse(c, errCode, "error binding", getErr)
	}

	return controllers.SuccessResponse(c, response.FromDomainList(data))
}
