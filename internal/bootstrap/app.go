package bootstrap

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-crud/config"
	v1 "github.com/xarick/golang-crud/internal/gateways/rest/v1"
	"github.com/xarick/golang-crud/internal/services"
)

type App struct {
	cfg    config.Application
	engine *gin.Engine
}

func New(cfg config.Application) *App {
	service := services.NewService(&cfg, CreateHTTPSClient())
	ctrl := v1.NewController(&cfg, service)

	r := gin.Default()
	engine := v1.NewRouter(r, ctrl)

	app := App{
		cfg:    cfg,
		engine: engine,
	}

	return &app
}

func (app *App) Run(ctx context.Context, cfg config.Application) {
	go func() {
		err := app.engine.Run(cfg.RunPort)
		if err != nil {
			log.Panic(err)
		}
	}()

	<-ctx.Done()
}

func CreateHTTPSClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &http.Client{Transport: tr}
}
