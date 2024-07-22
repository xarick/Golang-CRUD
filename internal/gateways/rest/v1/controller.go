package v1

import (
	"github.com/xarick/golang-crud/config"
	"github.com/xarick/golang-crud/internal/services"
)

type Controller struct {
	serv *services.Service
	cfg  *config.Application
}

func NewController(cfg *config.Application, serv *services.Service) *Controller {
	return &Controller{cfg: cfg, serv: serv}
}
