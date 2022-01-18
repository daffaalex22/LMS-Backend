package courses

import (
	_middleware "backend/app/middleware"
	"backend/business/course"
	"backend/controllers"
	"backend/controllers/courses/request"
	"backend/controllers/courses/response"
	"backend/helper/err"
	"log"

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

	listDomain := response.BatchesFromDomainList(data)
	return controllers.SuccessResponse(c, listDomain)
}

func (cl *CourseController) SearchCourses(c echo.Context) error {
	ctx := c.Request().Context()
	title := c.QueryParam("title")
	difficulty := c.QueryParam("difficulty")
	category := c.QueryParam("category")

	data, message := cl.CourseUsecase.SearchCourses(ctx, title, category, difficulty)

	if message != nil {
		codeErr := err.ErrorGetAllCourse(message)
		return controllers.ErrorResponse(c, codeErr, "error request", message)
	}

	listDomain := response.BatchesFromDomainList(data)
	return controllers.SuccessResponse(c, listDomain)
}

func (cl *CourseController) GetCourseById(c echo.Context) error {
	log.Println("INFO Service GetCourseById running")
	id := c.Param("courseId")

	ctx := c.Request().Context()
	data, message := cl.CourseUsecase.GetCourseById(ctx, id)

	if message != nil {
		codeErr, errorMessage := err.ErrorGetCourseById(message)
		return controllers.ErrorResponse(c, codeErr, errorMessage, message)

	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
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

func (cl *CourseController) Update(c echo.Context) error {
	req := request.UpdateRequest{}
	c.Bind(&req)

	id := c.Param("courseId")
	ctx := c.Request().Context()

	data, message := cl.CourseUsecase.Update(ctx, id, req.ToDomain())
	if message != nil {
		codeErr, errMessage := err.ErrorUpdateCourseCheck(message)
		return controllers.ErrorResponse(c, codeErr, errMessage, message)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (cl *CourseController) GetCourseByStudentId(c echo.Context) error {
	ctx := c.Request().Context()
	studentId := _middleware.GetIdFromJWT(c)

	data, message := cl.CourseUsecase.GetCourseByStudentId(ctx, uint(studentId))

	if message != nil {
		codeErr, errorMessage := err.ErrorGetCourseByStudentId(message)
		return controllers.ErrorResponse(c, codeErr, errorMessage, message)

	}
	return controllers.SuccessResponse(c, response.BatchesFromDomainList(data))
}

func (cl *CourseController) GetCourseByTeacherId(c echo.Context) error {
	ctx := c.Request().Context()
	studentId := _middleware.GetIdFromJWTtch(c)

	data, message := cl.CourseUsecase.GetCourseByTeacherId(ctx, uint(studentId))

	if message != nil {
		codeErr, errorMessage := err.ErrorGetCourseByTeacherId(message)
		return controllers.ErrorResponse(c, codeErr, errorMessage, message)
	}
	return controllers.SuccessResponse(c, response.BatchesFromDomainList(data))
}
