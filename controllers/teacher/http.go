package teacher

import (
	_middleware "backend/app/middleware"
	"backend/business/teacher"
	"backend/controllers"
	"backend/controllers/teacher/request"
	"backend/controllers/teacher/response"
	"net/http"

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
	err := c.Bind(&teacherLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	tch, err1 := controller.usecase.TeacherLogin(*teacherLogin.ToDomainLogin(), ctx)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err1)
	}
	return controllers.SuccessResponse(c, response.FromDomainLogin(tch))
}

func (controller *TeacherController) TeacherRegister(c echo.Context) error {
	ctx := c.Request().Context()
	reqRegist := request.TeacherRegister{}
	err := c.Bind(&reqRegist)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	tch, err1 := controller.usecase.TeacherRegister(reqRegist.ToDomainRegist(), ctx)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err1)
	}
	return controllers.SuccessResponse(c, response.FromDomainToRegist(tch))
}

func (controller *TeacherController) TeacherUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWTtch(c)
	var tchUpdate request.TeacherUpdate
	err := c.Bind(&tchUpdate)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	tch, err := controller.usecase.TeacherUpdate(ctx, *tchUpdate.ToDomainUpdate(), uint(id))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	return controllers.SuccessResponse(c, response.FromDomainToUpdate(tch))
}

func (controller *TeacherController) TeacherGetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	id := _middleware.GetIdFromJWTtch(c)
	tch, err := controller.usecase.TeacherGetProfile(ctx, uint(id))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomainProfile(tch))
}
