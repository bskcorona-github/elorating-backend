// backend/usecase/elo_usecase.go

package usecase

import (
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/repository"
)

type EloUsecase interface {
	// TODO: 必要なElo関連のユースケースメソッドをここに宣言する
}

type eloUsecase struct {
	eloRepository repository.EloRepository
}

func NewEloUsecase(eloRepository repository.EloRepository) EloUsecase {
	return &eloUsecase{
		eloRepository: eloRepository,
	}
}

// TODO: 必要なElo関連のユースケースメソッドを実装する
