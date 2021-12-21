package teacher

import (
	_middleware "backend/app/middleware"
	"backend/business/teacher"
	"backend/controllers"
	"backend/controllers/teacher/request"
	"backend/controllers/teacher/response"
	"backend/helper/err"

	"github.com/labstack/echo/v4"
)

type TeacherController struct {
	usecase teacher.TeacherUseCaseInterface
}

func NewTeacherController(tc teacher.TeacherUseCaseInterface) *TeacherController {
	return &TeacherController{
		usecase: tc,
	}
}

func (controller *TeacherController) TeacherLogin(c echo.Context) error {
	ctx := c.Request().Context()
	var teacherLogin request.TeacherLogin
	c.Bind(&teacherLogin)
	tch, result := controller.usecase.TeacherLogin(*teacherLogin.ToDomainLogin(), ctx)
	if result != nil {
		codeErr := err.ErrorTeacherLoginCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainLogin(tch))
}

func (controller *TeacherController) TeacherRegister(c echo.Context) error {
	ctx := c.Request().Context()
	reqRegist := request.TeacherRegister{}
	c.Bind(&reqRegist)
	tch, result := controller.usecase.TeacherRegister(reqRegist.ToDomainRegist(), ctx)
	if result != nil {
		codeErr := err.ErrorTeacherRegisterCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainToRegist(tch))
}

func (controller *TeacherController) TeacherUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWTtch(c)
	var tchUpdate request.TeacherUpdate
	c.Bind(&tchUpdate)
	tch, result := controller.usecase.TeacherUpdate(ctx, *tchUpdate.ToDomainUpdate(), uint(id))
	if result != nil {
		codeErr := err.ErrorTeacherUpdateCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainToUpdate(tch))
}

func (controller *TeacherController) TeacherGetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWTtch(c)
	tch, result := controller.usecase.TeacherGetProfile(ctx, uint(id))
	if result != nil {
		codeErr := err.ErrorTeacherProfileCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessResponse(c, response.FromDomainProfile(tch))
}
