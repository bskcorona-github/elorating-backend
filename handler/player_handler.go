package handlers

import (
	"net/http"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/usecase"
	"github.com/labstack/echo/v4"
)

type PlayerHandler struct {
	playerUsecase *usecase.PlayerUsecase
}

func NewPlayerHandler(playerUsecase *usecase.PlayerUsecase) *PlayerHandler {
	return &PlayerHandler{
		playerUsecase: playerUsecase,
	}
}

// プレイヤー全取得
func (ph *PlayerHandler) GetAllPlayers(c echo.Context) error {
	// ユースケースの呼び出し
	players, err := ph.playerUsecase.GetAllPlayers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, players)
}

// プレイヤー個別取得
func (ph *PlayerHandler) GetPlayerByID(c echo.Context) error {
	// パラメータからIDを取得
	playerID := c.Param("id")

	// ユースケースの呼び出し
	player, err := ph.playerUsecase.GetPlayerByID(playerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Player not found")
	}
	return c.JSON(http.StatusOK, player)
}

// プレイヤー作成
func (ph *PlayerHandler) CreatePlayer(c echo.Context) error {
	// リクエストボディからデータをパース
	var player models.Player
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// ユースケースの呼び出し
	newPlayer, err := ph.playerUsecase.CreatePlayer(&player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newPlayer)
}

// プレイヤー情報編集更新
func (ph *PlayerHandler) UpdatePlayer(c echo.Context) error {
	// パラメータからIDを取得
	playerID := c.Param("id")

	// リクエストボディからデータをパース
	var updatedPlayer *models.Player
	if err := c.Bind(&updatedPlayer); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// ユースケースの呼び出し
	updatedPlayer, err := ph.playerUsecase.UpdatePlayer(playerID, updatedPlayer)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Player not found")
	}
	return c.JSON(http.StatusOK, updatedPlayer) // ここでポインタ型を返す
}

// プレイヤー削除
func (ph *PlayerHandler) DeletePlayer(c echo.Context) error {
	// パラメータからIDを取得
	playerID := c.Param("id")

	// ユースケースの呼び出し
	err := ph.playerUsecase.DeletePlayer(playerID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Player not found")
	}
	return c.NoContent(http.StatusNoContent)
}
