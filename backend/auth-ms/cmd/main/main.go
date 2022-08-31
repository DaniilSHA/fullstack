package main

import (
	"context"
	auth_ms "fullstack/backend/auth-ms"
	"fullstack/backend/auth-ms/pkg/handler"
	"fullstack/backend/auth-ms/pkg/repository/mongodb"
	"fullstack/backend/auth-ms/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error initializing configs: %s", err.Error())
	//}

	//if err := godotenv.Load("./backend/auth-ms/.env"); err != nil {
	//	logrus.Fatalf("error loading env variables: %s", err.Error())
	//}

	mongoDBClient, err := mongodb.NewClient(context.Background(), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_AUTH_DB"))
	if err != nil {
		panic(err)
	}

	authRepository := mongodb.NewAuthMongo(mongoDBClient, os.Getenv("DB_COLLECTION"))
	services := service.NewAuthService(authRepository)
	handlers := handler.NewHandler(services)

	srv := new(auth_ms.Server)
	if err := srv.Run(os.Getenv("LISTEN_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("backend/auth-ms")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
