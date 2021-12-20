package routes

import (
	"backend/controllers/categories"
<<<<<<< HEAD
	"backend/controllers/courses"
=======
	enrollmentsController "backend/controllers/enrollments"
>>>>>>> 0152b61e00ded4a7614ad61e072888a0a690f1cb
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
<<<<<<< HEAD
	CategoryController categories.CategoriesController
	StudentController  studentController.StudentController
	TeacherController  teacherController.TeacherController
	CourseController   courses.CourseController
	JWTConfig          middleware.JWTConfig
	JWTConfigs         middleware.JWTConfig
=======
	CategoryController    categories.CategoriesController
	StudentController     studentController.StudentController
	TeacherController     teacherController.TeacherController
	EnrollmentsController enrollmentsController.EnrollmentsController
	JWTConfig             middleware.JWTConfig
	JWTConfigs            middleware.JWTConfig
>>>>>>> 0152b61e00ded4a7614ad61e072888a0a690f1cb
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

func (cl *RouteControllerList) CourseRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("api/v1/courses", cl.CourseController.Create)
}
