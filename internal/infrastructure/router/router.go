package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	{
		v1.POST("/estimation", estimationController.Estimate)
		v1.POST("/distance-categories", estimationController.SetDistanceCategories)
		v1.POST("/default-price", estimationController.SetDefaultPrice)
	}
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}
