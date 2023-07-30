package usecase

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/repository"
)

type PlayerUsecase struct {
	playerRepo *repository.PlayerRepository
}

func NewPlayerUsecase(playerRepo *repository.PlayerRepository) *PlayerUsecase {
	return &PlayerUsecase{
		playerRepo: playerRepo,
	}
}

// プレイヤー全取得
func (pu *PlayerUsecase) GetAllPlayers() ([]models.Player, error) {
	// リポジトリの呼び出し
	return pu.playerRepo.GetAllPlayers()
}

// プレイヤー個別取得
func (pu *PlayerUsecase) GetPlayerByID(playerID string) (*models.Player, error) {
	// リポジトリの呼び出し
	return pu.playerRepo.GetPlayerByID(playerID)
}

// プレイヤー作成
func (pu *PlayerUsecase) CreatePlayer(player *models.Player) (*models.Player, error) {
	// リポジトリの呼び出し
	return pu.playerRepo.CreatePlayer(player)
}

// プレイヤー情報編集更新
func (pu *PlayerUsecase) UpdatePlayer(playerID string, updatedPlayer *models.Player) (*models.Player, error) {
	// リポジトリの呼び出し
	return pu.playerRepo.UpdatePlayer(playerID, updatedPlayer)
}

// プレイヤー削除
func (pu *PlayerUsecase) DeletePlayer(playerID string) error {
	// リポジトリの呼び出し
	return pu.playerRepo.DeletePlayer(playerID)
}
