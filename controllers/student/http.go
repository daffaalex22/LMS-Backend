package student

import (
	_middleware "backend/app/middleware"
	"backend/business/student"
	"backend/controllers"
	"backend/controllers/student/request"
	"backend/controllers/student/response"
	"backend/helper/err"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	usecase student.StudentUseCaseInterface
}

func NewStudentController(sc student.StudentUseCaseInterface) *StudentController {
	return &StudentController{
		usecase: sc,
	}
}

func (controller *StudentController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var studentLogin request.StudentLogin
	c.Bind(&studentLogin)
	std, result := controller.usecase.Login(*studentLogin.ToDomainLogin(), ctx)
	if result != nil {
		codeErr := err.ErrorStudentLoginCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainLogin(std))
}

func (controller *StudentController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	reqRegist := request.StudentRegister{}
	c.Bind(&reqRegist)

	std, result := controller.usecase.Register(reqRegist.ToDomainRegist(), ctx)
	if result != nil {
		codeErr := err.ErrorStudentRegisterCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainToRegist(std))
}

func (controller *StudentController) StudentUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWT(c)
	var stdUpdate request.StudentUpdate
	c.Bind(&stdUpdate)
	std, result := controller.usecase.StudentUpdate(ctx, *stdUpdate.ToDomainUpdate(), uint(id))
	if result != nil {
		codeErr := err.ErrorStudentUpdateCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainToUpdate(std))
}

func (controller *StudentController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWT(c)
	std, result := controller.usecase.GetProfile(ctx, uint(id))
	if result != nil {
		codeErr := err.ErrorStudentUpdateCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainProfile(std))
}
