package handler

import (
	"net/http"

	"github.com/cool-service-go/repository"
	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

type UserHandler struct {
	repo *repository.Repository
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (us UserHandler) GetUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		users, err := us.repo.UserRepository.GetUsers(c.Request().Context())
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, H{"message": "ok", "data": users})
	}
}

func (us UserHandler) AddUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{"message": "ok"})
	}
}
