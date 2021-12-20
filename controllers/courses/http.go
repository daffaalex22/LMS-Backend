package courses

import (
	"backend/business/course"
	"backend/controllers"
	"backend/controllers/courses/request"
	"backend/controllers/courses/response"
	"backend/helpers/err"

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
		return controllers.NewErrorResponse(c, codeErr, message)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}
