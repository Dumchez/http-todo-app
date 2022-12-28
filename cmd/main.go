package main

import (
	"log"

	todo "github.com/Dumchez/http-todo-app"
	"github.com/Dumchez/http-todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Start("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while starting http server: %s", err.Error())
	}

}
