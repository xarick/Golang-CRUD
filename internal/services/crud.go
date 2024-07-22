package services

import (
	"context"
	"net/http"

	"github.com/xarick/golang-crud/config"
)

type CRUDService struct {
	client *http.Client
	cfg    *config.Application
}

func NewCRUDService(cfg *config.Application, client *http.Client) *CRUDService {
	return &CRUDService{cfg: cfg, client: client}
}

func (ch *CRUDService) GetAllFunc(ctx context.Context) (string, error) {
	return "Sent list", nil
}
