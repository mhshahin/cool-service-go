package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/model"
	"github.com/mhshahin/cool-service-go/service"
	"github.com/mhshahin/cool-service-go/utility/logger"
	"go.uber.org/zap"
)

type OpaMiddleware struct {
	service *service.Service
	cfg     *config.AppConfig
	logger  *zap.SugaredLogger
}

func NewOpaMiddleware(cfg *config.AppConfig, service *service.Service) *OpaMiddleware {
	return &OpaMiddleware{
		service: service,
		cfg:     cfg,
		logger:  logger.GetSugaredLogger(),
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
			om.logger.Errorw(
				"there was an error validating the access",
				"error", err,
			)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		if !result {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
