package handlers

import (
	"net/http"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/usecase"
	"github.com/labstack/echo/v4"
)

type EloHandler struct {
	eloUsecase usecase.EloUsecase
}

func NewEloHandler(eloUsecase usecase.EloUsecase) *EloHandler {
	return &EloHandler{
		eloUsecase: eloUsecase,
	}
}

// チーム分けAPIのハンドラ関数
func (h *EloHandler) TeamFormationHandler(c echo.Context) error {
	// リクエストから必要なデータを取得（例：選択した10人のプレイヤーIDなど）
	// 選択した10人のプレイヤーIDを元にチーム分けを行う
	teams, err := h.eloUsecase.TeamFormation(players)
	if err != nil {
		// エラーハンドリング
		return echo.NewHTTPError(http.StatusInternalServerError, "チーム分けに失敗しました")
	}

	// チーム分け結果をJSON形式でレスポンス
	return c.JSON(http.StatusOK, teams)
}

// ゲーム結果の入力APIのハンドラ関数
func (h *EloHandler) GameResultHandler(c echo.Context) error {
	// リクエストから必要なデータを取得（例：勝ったチームと負けたチームの情報など）
	// ゲーム結果を元にレートの変動を計算する
	newRatings, err := h.eloUsecase.CalculateRatings(gameResult)
	if err != nil {
		// エラーハンドリング
		return echo.NewHTTPError(http.StatusInternalServerError, "ゲーム結果の計算に失敗しました")
	}

	// レートの変動結果をJSON形式でレスポンス
	return c.JSON(http.StatusOK, newRatings)
}
