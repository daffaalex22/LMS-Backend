package types

import (
	"backend/business/types"
	"backend/controllers"
	"backend/controllers/types/response"
	"backend/helper/err"
	"fmt"

	"github.com/labstack/echo/v4"
)

type TypesController struct {
	TypesUsecase types.Usecase
}

func NewTypesController(typeUsecase types.Usecase) *TypesController {
	return &TypesController{
		TypesUsecase: typeUsecase,
	}
}

func (typeController TypesController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	data, getErr := typeController.TypesUsecase.GetAll(ctx)

	if getErr != nil {
		errCode := err.ErrorTypeCheck(getErr)
		fmt.Println(errCode)
		return controllers.ErrorResponse(c, errCode, "error binding", getErr)
	}

	return controllers.SuccessResponse(c, response.FromDomainList(data))
}
