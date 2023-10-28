package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/model"
	"github.com/mhshahin/cool-service-go/utility/logger"
	"go.uber.org/zap"
)

type JwtMiddleware struct {
	cfg    *config.AppConfig
	logger *zap.SugaredLogger
}

func NewJwtMiddleware(cfg *config.AppConfig) *JwtMiddleware {
	return &JwtMiddleware{
		cfg:    cfg,
		logger: logger.GetSugaredLogger(),
	}
}

func (jm JwtMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			jm.logger.Errorw(
				"empty token provided",
			)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jm.cfg.JWT.Secret), nil
			},
		)
		if err != nil {
			jm.logger.Errorw(
				"there was an error parsing the claims",
				"error", err,
			)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(*model.JwtCustomClaims)
		if !ok || !token.Valid {
			jm.logger.Infow(
				"invalid token provided",
				"claims", claims,
				"token", token,
			)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("authenticated", true)
		c.Set("role", claims.Role)

		jm.logger.Infow(
			"request authorized",
			"user", claims.Name,
			"role", claims.Role,
			"request_timestamp", time.Now().Format(time.RFC3339),
		)

		return next(c)
	}
}
