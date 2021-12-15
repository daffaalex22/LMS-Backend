package routes

import (
	"backend/controllers/categories"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	CategoryController categories.CategoriesController
}

func (cl *ControllerList) CategoriesRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("api/v1/categories")
	ev1.GET("/GetAll", cl.CategoryController.GetAll)
}
