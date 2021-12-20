package routes

import (
	"backend/controllers/categories"
	enrollmentsController "backend/controllers/enrollments"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	CategoryController    categories.CategoriesController
	StudentController     studentController.StudentController
	TeacherController     teacherController.TeacherController
	EnrollmentsController enrollmentsController.EnrollmentsController
	JWTConfig             middleware.JWTConfig
	JWTConfigs            middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(controller.JWTConfig)
	jwts := middleware.JWTWithConfig(controller.JWTConfigs)

	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("api/v1")
	//categories
	ev1.GET("/categories", controller.CategoryController.GetAll)

	//student
	ev1.POST("/students/login", controller.StudentController.Login)
	ev1.POST("/students/register", controller.StudentController.Register)
	ev1.GET("/students/profile", controller.StudentController.GetProfile, jwt)
	ev1.PUT("/students/profile", controller.StudentController.StudentUpdate, jwt)

	//teacher
	ev1.POST("/teacher/login", controller.TeacherController.TeacherLogin)
	ev1.POST("/teacher/register", controller.TeacherController.TeacherRegister)
	ev1.GET("/teacher/profile", controller.TeacherController.TeacherGetProfile, jwts)
	ev1.PUT("/teacher/profile", controller.TeacherController.TeacherUpdate, jwts)

	//enrollments
	ev1.GET("/enrollments", controller.EnrollmentsController.EnrollmentsGetAll)
}
