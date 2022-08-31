package main

import (
	core_ms "fullstack/backend/core-ms"
	handler "fullstack/backend/core-ms/pkg"
	"github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//IF YOU WANT START BY LOCALHOST WITH ENV
	if err := godotenv.Load("./backend/core-ms/.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	handlers := handler.NewHandler()
	srv := new(core_ms.Server)
	if err := srv.Run(os.Getenv("LISTEN_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server: %s", err)
	}
}
