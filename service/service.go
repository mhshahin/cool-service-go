package service

import (
	"github.com/mhshahin/cool-service-go/config"
	opaservice "github.com/mhshahin/cool-service-go/service/opa_service"
)

type Service struct {
	OpaService *opaservice.OpaService
}

func NewService(cfg *config.AppConfig) *Service {
	return &Service{
		OpaService: opaservice.NewOpaService(cfg),
	}
}
