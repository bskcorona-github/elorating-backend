package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "your_db_user"),
		DBPassword: getEnv("DB_PASSWORD", "your_db_password"),
		DBName:     getEnv("DB_NAME", "your_db_name"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func (c *Config) DBURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPassword)
}

func (c *Config) DBConnection() (*gorm.DB, error) {
	dsn := c.DBURI()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// データベースの接続数などの設定を行う場合はここに追加する

	return db, nil
}
