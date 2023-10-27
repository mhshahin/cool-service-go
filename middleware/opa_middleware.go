package middleware

import (
	"net/http"
	"strings"

	"github.com/cool-service-go/config"
	"github.com/cool-service-go/model"
	"github.com/cool-service-go/service"
	"github.com/labstack/echo/v4"
)

type OpaMiddleware struct {
	service *service.Service
	cfg     *config.AppConfig
}

func NewOpaMiddleware(cfg *config.AppConfig, service *service.Service) *OpaMiddleware {
	return &OpaMiddleware{
		service: service,
		cfg:     cfg,
	}
}

func (om *OpaMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !om.cfg.OPA.Enabled {
			return next(c)
		}

		path := strings.Split(c.Request().URL.Path, "/")

		newOpaRequest := model.OpaRequest{
			Input: model.Input{
				Path:            path[1:],
				Method:          c.Request().Method,
				IsAuthenticated: c.Get("authenticated").(bool),
				Role:            c.Get("role").(string),
			},
		}

		result, err := om.service.OpaService.Validate(newOpaRequest)
		if err != nil {
			return c.JSON(http.StatusForbidden, "unauthorized")
		}

		if !result {
			return c.JSON(http.StatusForbidden, "unauthorized")
		}

		return next(c)
	}
}
