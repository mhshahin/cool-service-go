package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/handler"
	"github.com/mhshahin/cool-service-go/middleware"
)

func InitRoutes(e *echo.Echo, handler *handler.Handler, middleware *middleware.Middleware) {
	api := e.Group("/api")

	api.GET("/users", handler.UserHandler.GetUsers(),
		middleware.JwtMiddleware.Middleware,
		middleware.OpaMiddleware.Middleware,
	)
	api.POST("/users", handler.UserHandler.AddUsers(),
		middleware.JwtMiddleware.Middleware,
		middleware.OpaMiddleware.Middleware,
	)

	authApi := api.Group("/auth")
	authApi.POST("/token", handler.AuthHandler.CreateToken())
}
