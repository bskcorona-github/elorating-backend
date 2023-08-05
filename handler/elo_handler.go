package handlers

import (
	"net/http"

	"log"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/usecase"
	"github.com/labstack/echo/v4"
)

// ELOHandler はelo関連のAPIハンドラーを持つ構造体です。
type ELOHandler struct {
	eloUsecase usecase.ELOUsecase
}

// NewELOHandler はELOHandlerのインスタンスを生成する関数です。
func NewELOHandler(eloUsecase usecase.ELOUsecase) *ELOHandler {
	return &ELOHandler{
		eloUsecase: eloUsecase,
	}
}

// SelectPlayersHandler はプレイヤー選択APIのハンドラー関数です。
func (h *ELOHandler) SelectPlayersHandler(c echo.Context) error {
	log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	var selectedPlayers []models.Player
	if err := c.Bind(&selectedPlayers); err != nil {
		log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", selectedPlayers)

		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// プレイヤー選択ロジックを呼び出す
	result, err := h.eloUsecase.TeamFormation(selectedPlayers)
	if err != nil {
		log.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", selectedPlayers)

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to form teams")
	}

	return c.JSON(http.StatusOK, result)
}

// GetTeamFormationHandler はチーム分け結果取得APIのハンドラー関数です。
func (h *ELOHandler) GetTeamFormationHandler(c echo.Context) error {
	// チーム分け結果取得ロジックを呼び出す
	result, err := h.eloUsecase.GetTeamFormation()
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get team formation")
	}

	return c.JSON(http.StatusOK, result)
}

// ReflectGameResultHandler はゲーム結果反映APIのハンドラー関数です。
func (h *ELOHandler) ReflectGameResultHandler(c echo.Context) error {
	var gameResult models.GameResult
	if err := c.Bind(&gameResult); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload")
	}

	// ゲーム結果反映ロジックを呼び出す
	if err := h.eloUsecase.CalculateRatings(gameResult); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to reflect game result")
	}

	return c.JSON(http.StatusOK, "game result reflected successfully")
}
