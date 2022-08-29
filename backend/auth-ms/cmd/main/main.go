package main

import (
	auth_ms "fullstack/backend/auth-ms"
	"fullstack/backend/auth-ms/internal/config"
	"fullstack/backend/auth-ms/pkg/handler"
	"log"
)

func main() {
	cfg := config.GetConfig()

	handlers := new(handler.Handler)

	srv := new(auth_ms.Server)
	if err := srv.Run(cfg.Listen.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error start server: %s", err)
	}
}
