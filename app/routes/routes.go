package routes

import (
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	StudentController studentController.StudentController
	TeacherController teacherController.TeacherController
	JWTConfig         middleware.JWTConfig
	JWTConfigs        middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(controller.JWTConfig)
	jwts := middleware.JWTWithConfig(controller.JWTConfigs)

	//student
	student := e.Group("/student")
	student.POST("/login", controller.StudentController.Login)
	student.POST("/register", controller.StudentController.Register)
	student.GET("/profile/", controller.StudentController.GetProfile, jwt)
	student.PUT("/profile/", controller.StudentController.StudentUpdate, jwt)

	//teacher
	teacher := e.Group("/teacher")
	teacher.POST("/login", controller.TeacherController.TeacherLogin)
	teacher.POST("/register", controller.TeacherController.TeacherRegister)
	teacher.GET("/profile/", controller.TeacherController.TeacherGetProfile, jwts)
	teacher.PUT("/profile/", controller.TeacherController.TeacherUpdate, jwts)
}
