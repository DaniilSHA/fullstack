package main

import (
	auth_ms "fullstack/backend/auth-ms"
	"fullstack/backend/auth-ms/internal/config"
	"fullstack/backend/auth-ms/pkg/handler"
	"fullstack/backend/auth-ms/pkg/repository"
	"fullstack/backend/auth-ms/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := config.GetConfig()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(auth_ms.Server)
	if err := srv.Run(cfg.Listen.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server: %s", err)
	}
}
