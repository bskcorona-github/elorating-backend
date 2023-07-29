package database

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/models"
	"gorm.io/gorm"
)

type PlayerDatabase interface {
	CreatePlayer(player *models.Player) error
}

type playerDatabase struct {
	db *gorm.DB
}

func NewPlayerDatabase(db *gorm.DB) PlayerDatabase {
	return &playerDatabase{
		db: db,
	}
}

func (d *playerDatabase) CreatePlayer(player *models.Player) error {
	if err := d.db.Create(player).Error; err != nil {
		return err
	}

	return nil
}
