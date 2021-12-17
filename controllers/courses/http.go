package courses

import (
	"backend/business/course"
	"backend/controllers"
	"backend/controllers/courses/request"
	"net/http"

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
	data, err := cl.CourseUsecase.Create(ctx, req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, data)
}
