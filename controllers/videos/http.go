package videos

import (
	"backend/business/videos"
	"backend/controllers"
	"backend/controllers/videos/request"
	"backend/controllers/videos/response"
	"backend/helper/err"
	"backend/helper/konversi"
	"fmt"

	"github.com/labstack/echo/v4"
)

type VideosController struct {
	rdsusecase videos.VideosUseCaseInterface
}

func NewVideosController(rdsc videos.VideosUseCaseInterface) *VideosController {
	return &VideosController{
		rdsusecase: rdsc,
	}
}

func (controller *VideosController) VideosAdd(c echo.Context) error {
	req := request.VideosAdd{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, result := controller.rdsusecase.VideosAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddVideosCheck(result)
		fmt.Printf("Error Usecase")
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *VideosController) VideosUpdate(c echo.Context) error {
	req := request.VideosUpdate{}
	c.Bind(&req)
	id := c.Param("id")
	konv, _ := konversi.StringToUint(id)
	ctx := c.Request().Context()
	data, result := controller.rdsusecase.VideosUpdate(ctx, req.ToDomain(), konv)

	if result != nil {
		fmt.Println("Update Usecase Error")
		codeErr := err.ErrorUpdateVideosCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *VideosController) VideosGetByModuleId(c echo.Context) error {
	moduleId := c.Param("moduleId")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(moduleId)
	if err1 != nil {
		codeErr := err.ErrorGetByModuleIdVideosCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	data, result := controller.rdsusecase.VideosGetByModuleId(ctx, konv)
	if result != nil {
		codeErr := err.ErrorGetByModuleIdVideosCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *VideosController) VideosDelete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		codeErr := err.ErrorDeleteModulesCheck(err1)
		return controllers.ErrorResponse(c, codeErr, "error param", err1)
	}
	result := controller.rdsusecase.VideosDelete(ctx, konv)
	if result != nil {
		codeErr := err.ErrorDeleteModulesCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.VideosResponse{Id: konv})
}
