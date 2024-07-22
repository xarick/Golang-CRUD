package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string
}

func LoadConfig() Application {
	cfg := Application{}

	cfg.GinMode = os.Getenv("GIN_MODE")
	if cfg.GinMode == gin.DebugMode || cfg.GinMode == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	cfg.RunPort = os.Getenv("RUN_PORT")
	return cfg
}
