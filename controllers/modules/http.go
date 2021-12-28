package modules

import (
	"backend/business/modules"
	"backend/controllers"
	"backend/controllers/modules/response"
	"backend/helper/err"

	"github.com/labstack/echo/v4"
)

type ModulesController struct {
	mdsusecase modules.ModulesRepoInterface
}

func NewModulesController(mdsc modules.ModulesUseCaseInterface) *ModulesController {
	return &ModulesController{
		mdsusecase: mdsc,
	}
}

func (controller *ModulesController) ModulesGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, result := controller.mdsusecase.ModulesGetAll(ctx)
	errCode := err.ErrorModulesCheck(result)
	if result != nil {
		return controllers.ErrorResponse(c, errCode, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}