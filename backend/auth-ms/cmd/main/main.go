package main

import (
	"context"
	auth_ms "fullstack/backend/auth-ms"
	"fullstack/backend/auth-ms/internal/config"
	"fullstack/backend/auth-ms/pkg/handler"
	"fullstack/backend/auth-ms/pkg/repository/mongodb"
	"fullstack/backend/auth-ms/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := config.GetConfig()

	mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.Auth_db)
	if err != nil {
		panic(err)
	}

	authRepository := mongodb.NewAuthMongo(mongoDBClient, cfg.MongoDB.Collection)
	services := service.NewAuthService(authRepository, cfg)
	handlers := handler.NewHandler(services)

	srv := new(auth_ms.Server)
	if err := srv.Run(cfg.Listen.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server: %s", err)
	}
}
