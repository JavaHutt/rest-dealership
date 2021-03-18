package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

func NewPostgresDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
		config.Timezone,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
