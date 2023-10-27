package middleware

import (
	"github.com/cool-service-go/config"
	"github.com/cool-service-go/service"
)

type Middleware struct {
	OpaMiddleware *OpaMiddleware
	JwtMiddleware *JwtMiddleware
}

func NewMiddleware(cfg *config.AppConfig, service *service.Service) *Middleware {
	return &Middleware{
		OpaMiddleware: NewOpaMiddleware(cfg, service),
		JwtMiddleware: NewJwtMiddleware(cfg),
	}
}
