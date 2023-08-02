package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dbURI string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// データベースの接続数などの設定を行う場合はここに追加する

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
