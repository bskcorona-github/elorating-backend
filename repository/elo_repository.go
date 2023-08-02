package repository

import (
	"fmt"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"gorm.io/gorm"
)

// EloRepository はELOに関するデータベース操作を抽象化するインターフェースです。
type EloRepository interface {
	SaveTeamFormationResult(result *models.TeamFormationResult) error
	UpdatePlayerRatings(players []*models.Player) error
	GetTeamFormationResult() (*models.TeamFormationResult, error)
}

// eloRepository はEloRepositoryの実装を持つ構造体です。
type eloRepository struct {
	db *gorm.DB
}

// NewEloRepository はeloRepositoryのインスタンスを生成する関数です。
func NewEloRepository(db *gorm.DB) EloRepository {
	return &eloRepository{
		db: db,
	}
}

// SaveTeamFormationResult はチーム分け結果をDBに保存するメソッドです。
func (r *eloRepository) SaveTeamFormationResult(result *models.TeamFormationResult) error {
	// チーム分け結果の保存ロジックを実装する
	// ここで result の情報を TeamFormationResults テーブルに保存します
	if err := r.db.Create(result).Error; err != nil {
		return fmt.Errorf("failed to save team formation result: %w", err)
	}
	return nil
}

// UpdatePlayerRatings はプレイヤーのレーティングを更新するメソッドです。
func (r *eloRepository) UpdatePlayerRatings(players []*models.Player) error {
	// プレイヤーのレーティング更新ロジックを実装する
	// ここで players の情報を players テーブルに反映します
	for _, player := range players {
		if err := r.db.Model(&models.Player{}).Where("name = ?", player.Name).Update("elo_rating", player.EloRating).Error; err != nil {
			return fmt.Errorf("failed to update player rating for %s: %w", player.Name, err)
		}
	}
	return nil
}

// GetTeamFormationResult はチーム分け結果をDBから取得するメソッドです。
func (r *eloRepository) GetTeamFormationResult() (*models.TeamFormationResult, error) {
	// チーム分け結果取得ロジックを実装する
	// ここで TeamFormationResults テーブルからチーム分け結果を取得します
	// 必要に応じてデータベースのクエリを実行して、チーム分け結果を取得してください
	// 仮に以下のような変数 result に取得したチーム分け結果を格納するものとします
	result := &models.TeamFormationResult{
		TeamA: []*models.Player{
			{Name: "player1", EloRating: 1500},
			{Name: "player2", EloRating: 1550},
			{Name: "player3", EloRating: 1450},
			{Name: "player4", EloRating: 1520},
			{Name: "player5", EloRating: 1480},
		},
		TeamB: []*models.Player{
			{Name: "player6", EloRating: 1530},
			{Name: "player7", EloRating: 1470},
			{Name: "player8", EloRating: 1510},
			{Name: "player9", EloRating: 1540},
			{Name: "player10", EloRating: 1460},
		},
	}
	return result, nil
}
