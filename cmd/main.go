package main

import (
	"log"

	todo "github.com/Dumchez/http-todo-app"
	"github.com/Dumchez/http-todo-app/pkg/handler"
	"github.com/Dumchez/http-todo-app/pkg/repository"
	"github.com/Dumchez/http-todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Start("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting http server: %s", err.Error())
	}

}
