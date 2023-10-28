package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/model"
)

type JwtMiddleware struct {
	cfg *config.AppConfig
}

func NewJwtMiddleware(cfg *config.AppConfig) *JwtMiddleware {
	return &JwtMiddleware{
		cfg: cfg,
	}
}

func (jm JwtMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jm.cfg.JWT.Secret), nil
			},
		)
		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*model.JwtCustomClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		c.Set("authenticated", true)
		c.Set("role", claims.Role)

		return next(c)
	}
}
