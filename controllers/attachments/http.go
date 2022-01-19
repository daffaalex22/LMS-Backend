package attachments

import (
	"backend/business/attachments"
	"backend/controllers"
	"backend/controllers/Attachments/request"
	"backend/controllers/Attachments/response"
	"backend/helper/err"
	"backend/helper/konversi"
	"fmt"

	"github.com/labstack/echo/v4"
)

type AttachmentsController struct {
	rdsusecase attachments.AttachmentsUseCaseInterface
}

func NewAttachmentsController(rdsc attachments.AttachmentsUseCaseInterface) *AttachmentsController {
	return &AttachmentsController{
		rdsusecase: rdsc,
	}
}

func (controller *AttachmentsController) AttachmentsAdd(c echo.Context) error {
	req := request.AttachmentsAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.rdsusecase.AttachmentsAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddAttachmentsCheck(result)
		fmt.Printf("Error Usecase")
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *AttachmentsController) AttachmentsUpdate(c echo.Context) error {
	req := request.AttachmentsUpdate{}
	c.Bind(&req)
	id := c.Param("id")
	konv, _ := konversi.StringToUint(id)
	ctx := c.Request().Context()
	data, result := controller.rdsusecase.AttachmentsUpdate(ctx, req.ToDomain(), konv)

	if result != nil {
		fmt.Println("Update Usecase Error")
		codeErr := err.ErrorUpdateAttachmentsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *AttachmentsController) AttachmentsGetById(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorGetByModuleIdAttachmentsCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.rdsusecase.AttachmentsGetByModuleId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByModuleIdAttachmentsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *AttachmentsController) AttachmentsDelete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorDeleteModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	result := controller.rdsusecase.AttachmentsDelete(ctx, konv)
	if result != nil {
		codeErr := err.ErrorDeleteModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.AttachmentsResponse{Id: konv})
}
