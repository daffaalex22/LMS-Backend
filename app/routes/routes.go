package routes

import (
	_middleware "backend/app/middleware"
	studentController "backend/controllers/student"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	StudentController studentController.StudentController
	JWTConfig         *_middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	//student
	student := e.Group("/student")
	student.POST("/login", controller.StudentController.Login)
	student.POST("/register", controller.StudentController.Register)
	student.GET("/profile/:id", controller.StudentController.GetProfile)
}
