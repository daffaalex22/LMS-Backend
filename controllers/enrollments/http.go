package enrollments

import (
	"backend/business/enrollments"
	"backend/controllers"
	"backend/controllers/enrollments/request"
	"backend/controllers/enrollments/response"
	"backend/helper/err"
	"backend/helper/konversi"

	"github.com/labstack/echo/v4"
)

type EnrollmentsController struct {
	elmusecase enrollments.EnrollmentsUseCaseInterface
}

func NewEnrollmentsController(elmc enrollments.EnrollmentsUseCaseInterface) *EnrollmentsController {
	return &EnrollmentsController{
		elmusecase: elmc,
	}
}

func (controller *EnrollmentsController) EnrollmentsGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, result := controller.elmusecase.EnrollmentGetAll(ctx)
	errCode := err.ErrorEnrollmentCheck(result)
	if result != nil {
		return controllers.ErrorResponse(c, errCode, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *EnrollmentsController) EnrollmentAdd(c echo.Context) error {
	req := request.EnrollAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.elmusecase.EnrollmentAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddEnrollCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *EnrollmentsController) EnrollUpdate(c echo.Context) error {
	req := request.EnrollUpdate{}
	c.Bind(&req)
	ctx := c.Request().Context()
	data, result := controller.elmusecase.EnrollUpdate(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorUpdateModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *EnrollmentsController) EnrollGetByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(courseId)
	if err1 != nil {
		codeErr := err.ErrorGetByCourseIdEnrollCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.elmusecase.EnrollGetByCourseId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByCourseIdEnrollCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}
