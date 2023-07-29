package repository

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/models"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	CreatePlayer(player *models.Player) (*models.Player, error)
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

func (r *playerRepository) CreatePlayer(player *models.Player) (*models.Player, error) {
	if err := r.db.Create(player).Error; err != nil {
		return nil, err
	}

	return player, nil
}
