package routes

import (
	_middleware "backend/app/middleware"
	"backend/controllers/categories"
	studentController "backend/controllers/student"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	CategoryController categories.CategoriesController
	StudentController  studentController.StudentController
	JWTConfig          *_middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {

	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("api/v1")
	//categories
	ev1.GET("/categories", controller.CategoryController.GetAll)

	//student
	ev1.POST("/students/login", controller.StudentController.Login)
	ev1.POST("/students/register", controller.StudentController.Register)
	ev1.GET("/students/profile/:id", controller.StudentController.GetProfile)
}
