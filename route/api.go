package route

import (
	"github.com/cool-service-go/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, handlers *handler.Handler) {
	api := e.Group("/api")

	api.GET("/users", handlers.UserHandler.GetUsers())
	api.POST("/users", handlers.UserHandler.AddUsers())
}
