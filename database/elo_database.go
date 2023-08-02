// backend/database/elo_database.go

package database

import (
	"gorm.io/gorm"
)

type EloDatabase struct {
	db *gorm.DB
}

func NewEloDatabase(db *gorm.DB) *EloDatabase {
	return &EloDatabase{
		db: db,
	}
}

// TODO: 必要なElo関連のデータベース処理を実装する
