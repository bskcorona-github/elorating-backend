package database

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"gorm.io/gorm"
)

type PlayerDatabase struct {
	db *gorm.DB
}

func NewPlayerDatabase(db *gorm.DB) *PlayerDatabase {
	return &PlayerDatabase{db: db}
}

func (pdb *PlayerDatabase) GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	err := pdb.db.Find(&players).Error
	return players, err
}

func (pdb *PlayerDatabase) GetPlayerByID(playerID string) (*models.Player, error) {
	var player models.Player
	err := pdb.db.Where("id = ?", playerID).First(&player).Error
	return &player, err
}

func (pdb *PlayerDatabase) CreatePlayer(player *models.Player) (*models.Player, error) {
	err := pdb.db.Create(player).Error
	return player, err
}

func (pdb *PlayerDatabase) UpdatePlayer(playerID string, updatedPlayer *models.Player) (*models.Player, error) {
	player, err := pdb.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	player.Name = updatedPlayer.Name
	player.EloRating = updatedPlayer.EloRating
	player.TotalGames = updatedPlayer.TotalGames
	player.TotalWins = updatedPlayer.TotalWins
	player.TotalLosses = updatedPlayer.TotalLosses

	err = pdb.db.Save(player).Error
	return player, err
}

func (pdb *PlayerDatabase) DeletePlayer(playerID string) error {
	player, err := pdb.GetPlayerByID(playerID)
	if err != nil {
		return err
	}

	err = pdb.db.Delete(player).Error
	return err
}
