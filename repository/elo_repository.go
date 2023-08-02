package repository

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (pr *PlayerRepository) GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	err := pr.db.Find(&players).Error
	return players, err
}

func (pr *PlayerRepository) GetPlayerByID(playerID string) (*models.Player, error) {
	var player models.Player
	err := pr.db.Where("id = ?", playerID).First(&player).Error
	return &player, err
}

func (pr *PlayerRepository) CreatePlayer(player *models.Player) (*models.Player, error) {
	err := pr.db.Create(player).Error
	return player, err
}

func (pr *PlayerRepository) UpdatePlayer(playerID string, updatedPlayer *models.Player) (*models.Player, error) {
	player, err := pr.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	player.Name = updatedPlayer.Name
	player.EloRating = updatedPlayer.EloRating
	player.TotalGames = updatedPlayer.TotalGames
	player.TotalWins = updatedPlayer.TotalWins
	player.TotalLosses = updatedPlayer.TotalLosses

	err = pr.db.Save(player).Error
	return player, err
}

func (pr *PlayerRepository) DeletePlayer(playerID string) error {
	player, err := pr.GetPlayerByID(playerID)
	if err != nil {
		return err
	}

	err = pr.db.Delete(player).Error
	return err
}

// 以下のメソッドを追加します

func (pr *PlayerRepository) GetPlayersByID(playerIDs []int) ([]models.Player, error) {
	var players []models.Player
	err := pr.db.Find(&players, playerIDs).Error
	return players, err
}
