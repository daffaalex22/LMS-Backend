package requests

import (
	"backend/business/requests"
	"backend/controllers"
	"backend/controllers/requests/request"
	"backend/controllers/requests/response"
	"backend/helper/err"
	"backend/helper/konversi"

	"github.com/labstack/echo/v4"
)

type RequestsController struct {
	requsecase requests.RequestsUseCaseInterface
}

func NewRequestsController(elmc requests.RequestsUseCaseInterface) *RequestsController {
	return &RequestsController{
		requsecase: elmc,
	}
}

func (controller *RequestsController) RequestsGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, result := controller.requsecase.RequestsGetAll(ctx)
	errCode := err.ErrorRequestsCheck(result)
	if result != nil {
		return controllers.ErrorResponse(c, errCode, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *RequestsController) RequestsAdd(c echo.Context) error {
	req := request.RequestsAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.requsecase.RequestsAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddRequestsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *RequestsController) RequestsUpdate(c echo.Context) error {
	req := request.RequestsUpdate{}
	c.Bind(&req)
	id := c.Param("id")
	konv, _ := konversi.StringToUint(id)
	ctx := c.Request().Context()
	data, result := controller.requsecase.RequestsUpdate(ctx, req.ToDomain(), konv)

	if result != nil {
		codeErr := err.ErrorUpdateModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *RequestsController) RequestsGetByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(courseId)
	if err1 != nil {
		codeErr := err.ErrorGetByCourseIdRequestsCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.requsecase.RequestsGetByCourseId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByCourseIdRequestsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}
