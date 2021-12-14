package student

import (
	"backend/business/student"
	"backend/controllers"
	"backend/controllers/student/request"
	"backend/controllers/student/response"
	"backend/helper/konversi"
	"net/http"

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
	err := c.Bind(&studentLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	std, err1 := controller.usecase.Login(*studentLogin.ToDomainLogin(), ctx)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err1)
	}
	return controllers.SuccessResponse(c, response.FromDomainLogin(std))
}

func (controller *StudentController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	reqRegist := request.StudentRegister{}
	err := c.Bind(&reqRegist)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	std, err1 := controller.usecase.Register(reqRegist.ToDomainRegist(), ctx)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err1)
	}
	return controllers.SuccessResponse(c, response.FromDomainToRegist(std))
}

func (controller *StudentController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	std, err := controller.usecase.GetProfile(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomainLogin(std))
}
