package repository

import (
	"fmt"
	"log"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"gorm.io/gorm"
)

// EloRepository はELOに関するデータベース操作を抽象化するインターフェースです。
type ELORepository interface {
	SaveTeamFormationResult(result *models.TeamFormationResult) (*models.TeamFormationResult, error)
	GetTeamFormationResult() (*models.TeamFormationResult, error)
	UpdatePlayerRatings(players []*models.Player) error
}

// eloRepository はEloRepositoryの実装を持つ構造体です。
type eloRepository struct {
	db *gorm.DB
}

// NewEloRepository はeloRepositoryのインスタンスを生成する関数です。
func NewELORepository(db *gorm.DB) *eloRepository {
	return &eloRepository{
		db: db,
	}
}

// SaveTeamFormationResult はチーム分け結果をDBに保存するメソッドです。
func (r *eloRepository) SaveTeamFormationResult(result *models.TeamFormationResult) (*models.TeamFormationResult, error) {
	// チーム分け結果の保存ロジックを実装する
	// ここで result の情報を TeamFormationResults テーブルに保存します
	log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", result)

	err := r.db.Create(result).Error
	return result, err
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
	// チーム分け結果を格納する変数
	result := &models.TeamFormationResult{}

	// チームAのプレイヤーIDを取得
	var teamAPlayerIDs string
	if err := r.db.Model(&models.TeamFormationResult{}).Pluck("team_a", &teamAPlayerIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to get team A player IDs: %w", err)
	}

	// チームBのプレイヤーIDを取得
	var teamBPlayerIDs string
	if err := r.db.Model(&models.TeamFormationResult{}).Pluck("team_b", &teamBPlayerIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to get team B player IDs: %w", err)
	}
	log.Println("1111111111111111111111111111111111111111", result)

	// チームAとチームBのプレイヤーIDをresultに設定
	result.TeamA = teamAPlayerIDs
	result.TeamB = teamBPlayerIDs

	return result, nil
}
