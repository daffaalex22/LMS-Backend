package enrollments

import (
	"backend/business/enrollments"
	"backend/controllers"
	"backend/controllers/enrollments/request"
	"backend/controllers/enrollments/response"
	"backend/helper/err"

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
