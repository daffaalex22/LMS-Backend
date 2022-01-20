package modules

import (
	"backend/business/modules"
	"backend/controllers"
	"backend/controllers/modules/request"
	"backend/controllers/modules/response"
	"backend/helper/err"
	"backend/helper/konversi"

	"github.com/labstack/echo/v4"
)

type ModulesController struct {
	mdsusecase modules.ModulesUseCaseInterface
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

func (controller *ModulesController) ModulesAdd(c echo.Context) error {
	req := request.ModulesAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.mdsusecase.ModulesAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ModulesController) ModulesUpdate(c echo.Context) error {
	req := request.ModulesUpdate{}
	c.Bind(&req)
	id := c.Param("id")
	konv, _ := konversi.StringToUint(id)
	ctx := c.Request().Context()
	data, result := controller.mdsusecase.ModulesUpdate(ctx, req.ToDomain(), konv)

	if result != nil {
		codeErr := err.ErrorUpdateModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ModulesController) ModulesGetByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(courseId)
	if err1 != nil {
		codeErr := err.ErrorGetByCourseIdModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.mdsusecase.ModulesGetByCourseId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByCourseIdModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *ModulesController) ModulesDelete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorDeleteModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	result := controller.mdsusecase.ModulesDelete(ctx, konv)
	if result != nil {
		codeErr := err.ErrorDeleteModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.ModulesResponse{Id: konv})
}

func (controller *ModulesController) ModulesGetById(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorGetModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	result, err2 := controller.mdsusecase.ModulesGetById(ctx, konv)
	if err2 != nil {
		codeErr := err.ErrorGetModulesCheck(err2)
		return controllers.ErrorResponse(c, codeErr, "error request", err2)
	}
	return controllers.SuccessResponse(c, response.FromDomain(result))
}
