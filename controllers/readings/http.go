package readings

import (
	"backend/business/readings"
	"backend/controllers"
	"backend/controllers/readings/request"
	"backend/controllers/readings/response"
	"backend/helper/err"
	"backend/helper/konversi"
	"fmt"

	"github.com/labstack/echo/v4"
)

type ReadingsController struct {
	rdsusecase readings.ReadingsUseCaseInterface
}

func NewReadingsController(rdsc readings.ReadingsUseCaseInterface) *ReadingsController {
	return &ReadingsController{
		rdsusecase: rdsc,
	}
}

func (controller *ReadingsController) ReadingsAdd(c echo.Context) error {
	req := request.ReadingsAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.rdsusecase.ReadingsAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddReadingsCheck(result)
		fmt.Printf("Error Usecase")
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ReadingsController) ReadingsUpdate(c echo.Context) error {
	req := request.ReadingsUpdate{}
	c.Bind(&req)
	id := c.Param("id")
	konv, _ := konversi.StringToUint(id)
	ctx := c.Request().Context()
	data, result := controller.rdsusecase.ReadingsUpdate(ctx, req.ToDomain(), konv)

	if result != nil {
		fmt.Println("Update Usecase Error")
		codeErr := err.ErrorUpdateReadingsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ReadingsController) ReadingsGetByModuleId(c echo.Context) error {
	moduleId := c.Param("moduleId")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(moduleId)
	if err1 != nil {
		codeErr := err.ErrorGetByModuleIdReadingsCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.rdsusecase.ReadingsGetByModuleId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByModuleIdReadingsCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *ReadingsController) ReadingsDelete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorDeleteModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	result := controller.rdsusecase.ReadingsDelete(ctx, konv)
	if result != nil {
		codeErr := err.ErrorDeleteModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.ReadingsResponse{Id: konv})
}
