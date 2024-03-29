package usecase

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/repository"
)

// ELORepository はelo関連のデータベース操作を行うインターフェースです。
type ELORepository interface {
	SaveTeamFormationResult(result *models.TeamFormationResult) (*models.TeamFormationResult, error)
	GetTeamFormationResult() (*models.TeamFormationResult, error)
	UpdatePlayerRatings(players []*models.Player) error
}

// eloUsecase はELOUsecaseの実装を持つ構造体です。
type ELOUsecase struct {
	playerRepository repository.PlayerRepository
	eloRepository    ELORepository
}

func sliceToString(slice []uint) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprint(v)
	}
	return strings.Join(strSlice, ",")
}

// NewELOUsecase はeloUsecaseのインスタンスを生成する関数です。
func NewELOUsecase(playerRepository repository.PlayerRepository, eloRepository repository.ELORepository) *ELOUsecase {
	return &ELOUsecase{
		playerRepository: playerRepository,
		eloRepository:    eloRepository,
	}
}

// TeamFormation はプレイヤーをチームに分けるロジックを実行するメソッドです。
func (u *ELOUsecase) TeamFormation(selectedPlayers []models.Player) (*models.TeamFormationResult, error) {
	if len(selectedPlayers) != 10 {
		return nil, errors.New("selectedPlayers must contain exactly 10 players")
	}

	// プレイヤーをEloレーティングでソートする
	sort.Slice(selectedPlayers, func(i, j int) bool {
		return selectedPlayers[i].EloRating > selectedPlayers[j].EloRating
	})

	// チームAとチームBのプレイヤーのIDを格納するスライス
	var teamAIDs []uint
	var teamBIDs []uint

	// プレイヤーを交互にチームAとチームBに追加していく
	for i, player := range selectedPlayers {
		if i%2 == 0 {
			teamAIDs = append(teamAIDs, player.ID)
		} else {
			teamBIDs = append(teamBIDs, player.ID)
		}
	}
	// チームAとチームBのプレイヤーのIDを文字列としてJSON形式に変換
	teamAString := sliceToString(teamAIDs)
	teamBString := sliceToString(teamBIDs)

	result := &models.TeamFormationResult{
		TeamA: teamAString,
		TeamB: teamBString,
	}
	log.Println("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm", result)

	// チーム分け結果をDBに保存
	result, err := u.eloRepository.SaveTeamFormationResult(result)
	if err != nil {
		log.Println("2222222222222222222222222222222222222222222222", result)

		return nil, fmt.Errorf("failed to save team formation result: %w", err)
	}

	return result, nil
}

// GetTeamFormation はチーム分け結果を取得するメソッドです。
func (u *ELOUsecase) GetTeamFormation() (*models.TeamFormationResult, error) {
	result, err := u.eloRepository.GetTeamFormationResult()
	if err != nil {
		return nil, fmt.Errorf("failed to get team formation result: %w", err)
	}
	return result, nil
}

// CalculateRatings はゲーム結果に応じてプレイヤーのレーティングを計算するメソッドです。
func (u *ELOUsecase) CalculateRatings(gameResult models.GameResult) error {
	// ゲーム結果計算ロジックを実装する

	// 勝ったチームのプレイヤーのレーティングを上げる
	for _, player := range gameResult.WinningTeam {
		player.EloRating += 10 // 仮の実装（勝った場合のレーティングの上昇値）
	}

	// 負けたチームのプレイヤーのレーティングを下げる
	for _, player := range gameResult.LosingTeam {
		player.EloRating -= 10 // 仮の実装（負けた場合のレーティングの減少値）
	}

	// プレイヤーのレーティングをDBに反映
	allPlayers := append(gameResult.WinningTeam, gameResult.LosingTeam...)
	err := u.eloRepository.UpdatePlayerRatings(allPlayers)
	if err != nil {
		return fmt.Errorf("failed to update player ratings: %w", err)
	}

	return nil
}
