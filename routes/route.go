package routes

import (
	"prakerja12/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/news", controllers.CreateNewsController)
	e.GET("/news", controllers.GetNewsController)
}