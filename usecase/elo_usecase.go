// backend/usecase/elo_usecase.go

package usecase

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/domain"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
)

type EloUsecase interface {
	TeamFormation(players []*models.Player) ([]*models.Team, error)
	CalculateRatings(gameResult *models.GameResult) (map[string]int, error)
}

type eloUsecase struct {
	eloRepository domain.EloRepository
}

func NewEloUsecase(eloRepo domain.EloRepository) EloUsecase {
	return &eloUsecase{
		eloRepository: eloRepo,
	}
}

func (u *eloUsecase) TeamFormation(players []*models.Player) ([]*models.Team, error) {
	// チーム分けのロジックを実装する
	// レーティングシステムなどを使用して、適切にチームを分ける
	// ここではeloRepositoryを使ってチーム分けを行う
	teams, err := u.eloRepository.TeamFormation(players)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (u *eloUsecase) CalculateRatings(gameResult *models.GameResult) (map[string]int, error) {
	// ゲーム結果の計算処理を実装する
	// 勝利したチームと敗北したチームのレートを更新するロジックを実装する
	// ここではeloRepositoryを使ってレート計算を行う
	ratings, err := u.eloRepository.CalculateRatings(gameResult)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
