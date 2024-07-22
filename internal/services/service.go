package services

import (
	"net/http"

	"github.com/xarick/golang-crud/config"
)

type Service struct {
	cfg     *config.Application
	CRUDSer *CRUDService
}

func NewService(cfg *config.Application, client *http.Client) *Service {
	return &Service{
		cfg:     cfg,
		CRUDSer: NewCRUDService(cfg, client),
	}
}
