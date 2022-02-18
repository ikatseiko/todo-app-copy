package main

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/ikatseiko/todo-app-copy/pkg/handler"
	"github.com/ikatseiko/todo-app-copy/pkg/repository"
	"github.com/ikatseiko/todo-app-copy/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initconfig(); err != nil {
		log.Fatalf("error initializing config: %s", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initconfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
