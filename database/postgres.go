// backend/database/postgres.go
package database

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
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

func (d *Database) CreatePlayer(player *models.Player) error {
	// TODO: Implement the logic to create a new player record in the database
	// This involves inserting the player information into the "players" table.
	return nil
}
