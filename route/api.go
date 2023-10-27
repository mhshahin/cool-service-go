package route

import (
	"github.com/cool-service-go/handler"
	"github.com/cool-service-go/middleware"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, handler *handler.Handler, middleware *middleware.Middleware) {
	api := e.Group("/api")

	api.GET("/users", handler.UserHandler.GetUsers(),
		middleware.JwtMiddleware.Middleware,
		middleware.OpaMiddleware.Middleware,
	)
	api.POST("/users", handler.UserHandler.AddUsers())

	authApi := api.Group("/auth")
	authApi.POST("/token", handler.AuthHandler.CreateToken())
}
