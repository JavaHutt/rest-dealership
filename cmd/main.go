package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/JavaHutt/rest-dealership/internal/action"
	"github.com/JavaHutt/rest-dealership/internal/handler"
	"github.com/JavaHutt/rest-dealership/internal/model"
	"github.com/JavaHutt/rest-dealership/internal/repository"
	"github.com/JavaHutt/rest-dealership/internal/service"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(new(model.Vehicle))
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed to get environment variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Timezone: viper.GetString("db.timezone"),
	})
	if err != nil {
		logrus.Fatalf("Failed to connect to db: %s", err.Error())
	}
	if err = autoMigrate(db); err != nil {
		logrus.Fatalf("Failed to perform migrations: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(action.Server)
	port := viper.GetString("port")
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while starting http server: %s", err.Error())
		}
	}()
	logrus.Printf("Rest dealership server started on port %s...", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Shutting down rest dealership server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error occured while shutting down server: %s", err.Error())
	}
}
