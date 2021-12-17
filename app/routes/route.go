package routes

import (
	"backend/controllers/categories"
	"backend/controllers/courses"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	CategoryController categories.CategoriesController
	CourseController   courses.CourseController
}

func (cl *ControllerList) CategoriesRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("api/v1")
	ev1.GET("/categories", cl.CategoryController.GetAll)
}

func (cl *ControllerList) CourseRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("api/v1/courses", cl.CourseController.Create)
}
