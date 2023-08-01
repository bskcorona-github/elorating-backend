// backend/repository/elo_repository.go

package repository

import (
	"gorm.io/gorm"
)

type EloRepository interface {
	// TODO: 必要なElo関連のリポジトリメソッドをここに宣言する
}

type eloRepository struct {
	db *gorm.DB
}

func NewEloRepository(db *gorm.DB) EloRepository {
	return &eloRepository{
		db: db,
	}
}

// TODO: 必要なElo関連のリポジトリメソッドを実装する
