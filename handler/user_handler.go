package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhshahin/cool-service-go/model"
	"github.com/mhshahin/cool-service-go/repository"
	"github.com/mhshahin/cool-service-go/utility/logger"
	"go.uber.org/zap"
)

type UserHandler struct {
	repo   *repository.Repository
	logger *zap.SugaredLogger
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{
		repo:   repo,
		logger: logger.GetSugaredLogger(),
	}
}

func (us UserHandler) GetUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		// TODO: implement pagination
		users, err := us.repo.UserRepository.GetUsers(c.Request().Context())
		if err != nil {
			us.logger.Errorw(
				"there was an error retrieving users from database",
				"error", err,
			)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, H{"message": "ok", "data": users})
	}
}

func (us UserHandler) AddUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		var users []*model.User

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			us.logger.Errorw(
				"there was an error reading the request body",
				"error", err,
			)
			return echo.ErrInternalServerError
		}

		err = json.Unmarshal(body, &users)
		if err != nil {
			us.logger.Errorw(
				"there was an error in unmarshal body",
				"error", err,
			)
			return echo.ErrInternalServerError
		}

		err = us.repo.UserRepository.AddUsers(c.Request().Context(), users)
		if err != nil {
			us.logger.Errorw(
				"there was an error in adding users to database",
				"error", err,
			)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, H{"message": "ok", "data": fmt.Sprintf("Successfully added %d users.", len(users))})
	}
}
