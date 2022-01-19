package routes

import (
	"backend/controllers/categories"
	"backend/controllers/courses"
	difficultiesController "backend/controllers/difficulties"
	enrollmentsController "backend/controllers/enrollments"
	modulesController "backend/controllers/modules"
	readingsController "backend/controllers/readings"
	requestsController "backend/controllers/requests"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"
	typesController "backend/controllers/types"
	videosController "backend/controllers/videos"

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
	RequestsController    requestsController.RequestsController
	ModulesController     modulesController.ModulesController
	ReadingsController    readingsController.ReadingsController
	VideosController      videosController.VideosController
	DifficultyController  difficultiesController.DifficultiesController
	TypeController        typesController.TypesController
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(controller.JWTConfig)
	jwts := middleware.JWTWithConfig(controller.JWTConfigs)

	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("api/v1")

	//categories
	ev1.GET("/categories", controller.CategoryController.GetAll)

	//difficulties
	ev1.GET("/difficulties", controller.DifficultyController.GetAll)

	//types
	ev1.GET("/types", controller.TypeController.GetAll)

	//student
	ev1.POST("/students/login", controller.StudentController.Login)
	ev1.POST("/students/register", controller.StudentController.Register)
	ev1.GET("/students/profile", controller.StudentController.GetProfile, jwt)
	ev1.PUT("/students/profile", controller.StudentController.StudentUpdate, jwt)

	//teacher
	ev1.POST("/teachers/login", controller.TeacherController.TeacherLogin)
	ev1.POST("/teachers/register", controller.TeacherController.TeacherRegister)
	ev1.GET("/teachers/profile", controller.TeacherController.TeacherGetProfile, jwts)
	ev1.PUT("/teachers/profile", controller.TeacherController.TeacherUpdate, jwts)

	//enrollments
	ev1.GET("/enrollments", controller.EnrollmentsController.EnrollmentsGetAll)
	ev1.POST("/enrollments", controller.EnrollmentsController.EnrollmentAdd)
	ev1.PUT("/enrollments", controller.EnrollmentsController.EnrollUpdate)

	//requests
	ev1.GET("/requests", controller.RequestsController.RequestsGetAll)
	ev1.POST("/requests", controller.RequestsController.RequestsAdd)
	ev1.GET("/requests/:id", controller.RequestsController.RequestGetById)
	ev1.PUT("/requests/:id", controller.RequestsController.RequestsUpdate)

	ev1.GET("/students/requests", controller.RequestsController.RequestsGetByStudentId, jwt)
	ev1.GET("/teachers/requests", controller.RequestsController.RequestsGetByTeacherId, jwts)

	//modules
	ev1.GET("/modules", controller.ModulesController.ModulesGetAll)
	ev1.POST("/modules", controller.ModulesController.ModulesAdd)
	ev1.PUT("/modules/:id", controller.ModulesController.ModulesUpdate)
	ev1.DELETE("/modules/:id", controller.ModulesController.ModulesDelete)

	//readings
	ev1.POST("/readings", controller.ReadingsController.ReadingsAdd)
	ev1.PUT("/readings/:id", controller.ReadingsController.ReadingsUpdate)
	ev1.DELETE("/readings/:id", controller.ReadingsController.ReadingsDelete)

	//videos
	ev1.POST("/videos", controller.VideosController.VideosAdd)
	ev1.PUT("/videos/:id", controller.VideosController.VideosUpdate)
	ev1.DELETE("/videos/:id", controller.VideosController.VideosDelete)

	//course
	ev1.GET("/courses/:courseId/modules", controller.ModulesController.ModulesGetByCourseId)
	ev1.GET("/courses/:courseId/enrollments", controller.EnrollmentsController.EnrollGetByCourseId)
	ev1.GET("/modules/:moduleId/readings", controller.ReadingsController.ReadingsGetByModuleId)
	ev1.GET("/modules/:moduleId/videos", controller.VideosController.VideosGetByModuleId)
	ev1.POST("/courses", controller.CourseController.Create)
	ev1.GET("/courses", controller.CourseController.GetAll)
	ev1.GET("/courses/search", controller.CourseController.SearchCourses)
	ev1.GET("/courses/:courseId", controller.CourseController.GetCourseById)
	ev1.PUT("/courses/:courseId", controller.CourseController.Update)
	ev1.DELETE("/courses/:courseId", controller.CourseController.Delete)

	ev1.GET("/students/courses", controller.CourseController.GetCourseByStudentId, jwt)
	ev1.GET("/teachers/courses", controller.CourseController.GetCourseByTeacherId, jwts)
}
