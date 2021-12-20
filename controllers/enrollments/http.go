package enrollments

import (
	"backend/business/enrollments"
	"backend/controllers"
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
	data, err1 := controller.elmusecase.EnrollmentGetAll(ctx)
	errCode := err.ErrorEnrollmentCheck(err1)
	if err1 != nil {
		return controllers.ErrorResponse(c, errCode, "error binding", err1)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}
