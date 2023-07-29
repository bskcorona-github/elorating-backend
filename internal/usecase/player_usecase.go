package usecase

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/repository"
)

type PlayerUsecase interface {
	CreatePlayer(player *models.Player) (*models.Player, error)
}

type playerUsecase struct {
	playerRepository repository.PlayerRepository
}

func NewPlayerUsecase(playerRepository repository.PlayerRepository) PlayerUsecase {
	return &playerUsecase{
		playerRepository: playerRepository,
	}
}

func (u *playerUsecase) CreatePlayer(player *models.Player) (*models.Player, error) {
	// プレイヤー情報のバリデーションなどのロジックを追加する場合はここに実装する

	createdPlayer, err := u.playerRepository.CreatePlayer(player)
	if err != nil {
		return nil, err
	}

	return createdPlayer, nil
}
