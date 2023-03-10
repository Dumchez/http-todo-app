package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	todo "github.com/Dumchez/http-todo-app"
	"github.com/Dumchez/http-todo-app/pkg/handler"
	"github.com/Dumchez/http-todo-app/pkg/repository"
	"github.com/Dumchez/http-todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title TODO APP API
// @version 1.0
// @description Swagger API Documentation for TODO APP

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initializing a database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while starting http server: %s", err.Error())
		}
	}()

	logrus.Print("Todo App Started Successfully!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Todo App Shutting Down!")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutdown: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured when closing db connection: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
