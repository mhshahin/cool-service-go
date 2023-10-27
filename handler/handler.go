package handler

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler() *Handler {
	return &Handler{
		UserHandler: NewUserHandler(),
	}
}
