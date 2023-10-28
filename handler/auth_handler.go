package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/model"
)

type AuthHandler struct {
	cfg *config.AppConfig
}

func NewAuthHandler(cfg *config.AppConfig) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (ah AuthHandler) CreateToken() func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.ErrInternalServerError
		}

		newJwtReq := model.JwtRequest{}
		err = json.Unmarshal(reqBody, &newJwtReq)
		if err != nil {
			return echo.ErrInternalServerError
		}

		claims := &model.JwtCustomClaims{
			Name:  newJwtReq.Name,
			Email: newJwtReq.Email,
			Role:  newJwtReq.Role,

			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(ah.cfg.JWT.ExpirationDuration)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(ah.cfg.JWT.Secret))
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, H{"message": "ok", "token": t})
	}
}
