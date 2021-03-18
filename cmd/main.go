package main

import (
	"log"

	"github.com/JavaHutt/rest-dealership/internal/action"
	"github.com/JavaHutt/rest-dealership/internal/handler"
	"github.com/JavaHutt/rest-dealership/internal/repository"
	"github.com/JavaHutt/rest-dealership/internal/service"

	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Timezone: viper.GetString("db.timezone"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(action.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while starting http server: %s", err.Error())
	}
}
