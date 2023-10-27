package handler

import (
	"github.com/cool-service-go/repository"
	"github.com/cool-service-go/service"
)

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(repo *repository.Repository, service *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(repo),
	}
}
