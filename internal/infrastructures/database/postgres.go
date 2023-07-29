package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "your_db_user"
	DBPassword = "your_db_password"
	DBName     = "your_db_name"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBName, DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// データベースの接続数などの設定を行う場合はここに追加する

	return db, nil
}
