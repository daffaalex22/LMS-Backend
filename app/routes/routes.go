package routes

import (
	"backend/controllers/categories"
	"backend/controllers/courses"
	enrollmentsController "backend/controllers/enrollments"
	modulesController "backend/controllers/modules"
	readingsController "backend/controllers/readings"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	CategoryController    categories.CategoriesController
	StudentController     studentController.StudentController
	TeacherController     teacherController.TeacherController
	CourseController      courses.CourseController
	JWTConfig             middleware.JWTConfig
	JWTConfigs            middleware.JWTConfig
	EnrollmentsController enrollmentsController.EnrollmentsController
	ModulesController     modulesController.ModulesController
	ReadingsController    readingsController.ReadingsController
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
	ev1.GET("/students/courses", controller.CourseController.GetCourseByStudentId, jwt)

	//teacher
	ev1.POST("/teacher/login", controller.TeacherController.TeacherLogin)
	ev1.POST("/teacher/register", controller.TeacherController.TeacherRegister)
	ev1.GET("/teacher/profile", controller.TeacherController.TeacherGetProfile, jwts)
	ev1.PUT("/teacher/profile", controller.TeacherController.TeacherUpdate, jwts)

	//enrollments
	ev1.GET("/enrollments", controller.EnrollmentsController.EnrollmentsGetAll)
	ev1.POST("/enrollments", controller.EnrollmentsController.EnrollmentAdd)
	ev1.PUT("/enrollments", controller.EnrollmentsController.EnrollUpdate)

	//modules
	ev1.GET("/modules", controller.ModulesController.ModulesGetAll)
	ev1.POST("/modules", controller.ModulesController.ModulesAdd)
	ev1.PUT("/modules/:id", controller.ModulesController.ModulesUpdate)
	ev1.DELETE("/modules/:id", controller.ModulesController.ModulesDelete)

	//readings
	ev1.POST("/readings", controller.ReadingsController.ReadingsAdd)
	ev1.PUT("/readings/:id", controller.ReadingsController.ReadingsUpdate)
	ev1.DELETE("/readings/:id", controller.ReadingsController.ReadingsDelete)

	//course
	ev1.GET("/courses/:courseId/modules", controller.ModulesController.ModulesGetByCourseId)
	ev1.GET("/courses/:courseId/enrollments", controller.EnrollmentsController.EnrollGetByCourseId)
	ev1.GET("/modules/:moduleId/readings", controller.ReadingsController.ReadingsGetByModuleId)
	ev1.POST("/courses", controller.CourseController.Create)
	ev1.GET("/courses", controller.CourseController.GetAll)
	ev1.GET("/courses/:courseId", controller.CourseController.GetCourseById)
	ev1.PUT("/courses/:courseId", controller.CourseController.Update)
	ev1.DELETE("/courses/:courseId", controller.CourseController.Delete)
}
