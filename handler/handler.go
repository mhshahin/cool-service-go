package handler

import "github.com/cool-service-go/repository"

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(repo),
	}
}
