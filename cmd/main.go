package main

import (
	"log"

	"github.com/JavaHutt/rest-dealership/internal/action"
	"github.com/JavaHutt/rest-dealership/internal/handler"
	"github.com/JavaHutt/rest-dealership/internal/repository"
	"github.com/JavaHutt/rest-dealership/internal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(action.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting http server: %s", err.Error())
	}
}
