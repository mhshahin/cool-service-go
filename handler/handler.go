package handler

import (
	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/repository"
	"github.com/mhshahin/cool-service-go/service"
)

type H map[string]interface{}

type Handler struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandler(cfg *config.AppConfig, repo *repository.Repository, service *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(repo),
		AuthHandler: NewAuthHandler(cfg),
	}
}
