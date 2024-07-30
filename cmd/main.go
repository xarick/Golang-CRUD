package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-crud/config"
	"github.com/xarick/golang-crud/internal/bootstrap"
)

func main() {
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGTERM)

	f, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	fEr, err := os.OpenFile("gin-error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer fEr.Close()

	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(fEr, os.Stdout)

	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := config.LoadConfig()
	app := bootstrap.New(cfg)

	// created and write [] to users.json
	fJson, err := os.OpenFile(cfg.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // os.O_RDWR
	if err != nil {
		log.Panic(err)
	}
	defer fJson.Close()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		OSCall := <-quitSignal
		log.Printf("System Call: %+v", OSCall)
		cancel()
	}()

	app.Run(ctx, cfg)
}
