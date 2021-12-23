package courses

import (
	"backend/business/course"
	"backend/controllers"
	"backend/controllers/courses/request"
	"backend/controllers/courses/response"
	"backend/helper/err"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	CourseUsecase course.Usecase
}

func NewCourseController(courseUsecase course.Usecase) *CourseController {
	return &CourseController{
		CourseUsecase: courseUsecase,
	}
}

func (cl *CourseController) Create(c echo.Context) error {
	req := request.AddRequest{}
	c.Bind(&req)

	ctx := c.Request().Context()
	data, message := cl.CourseUsecase.Create(ctx, req.ToDomain())

	if message != nil {
		codeErr := err.ErrorCreateCourse(message)
		return controllers.ErrorResponse(c, codeErr, "error request", message)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (cl *CourseController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, message := cl.CourseUsecase.GetAll(ctx)

	if message != nil {
		codeErr := err.ErrorGetAllCourse(message)
		return controllers.ErrorResponse(c, codeErr, "error request", message)
	}

	listDomain := response.FromDomainList(data)
	return controllers.SuccessResponse(c, listDomain)
}

func (cl *CourseController) Delete(c echo.Context) error {
	id := c.Param("courseId")

	ctx := c.Request().Context()
	data, message := cl.CourseUsecase.Delete(ctx, id)

	if message != nil {
		codeErr, errorMessage := err.ErrorDeleteCourse(message)
		return controllers.ErrorResponse(c, codeErr, errorMessage, message)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
