package blog

import (
	"fmt"
	"log"
	"time"
)

func Printf(format, path string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	t := time.Now().Format(time.RFC3339)
	log.Printf(`{"time": "%s", "level": "printf", "path": "%s", "msg": "%s"}`, t, path, msg)
}

func Info(path, msg string) {
	t := time.Now().Format(time.RFC3339)
	log.Printf(`{"time": "%s", "level": "info", "path": "%s", "msg": "%s"}`, t, path, msg)
}

func Error(path, msg string) {
	t := time.Now().Format(time.RFC3339)
	log.Printf(`{"time": "%s", "level": "error", "path": "%s", "msg": "%s"}`, t, path, msg)
}
