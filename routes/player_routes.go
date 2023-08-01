package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

func RegisterPlayerRoutes(e *echo.Echo, playerHandler handlers.PlayerHandler) {
	e.GET("/api/players", playerHandler.GetAllPlayers)     //プレイヤー全取得
	e.GET("/api/players/:id", playerHandler.GetPlayerByID) //プレイヤー個別取得
	// e.POST("/api/players", playerHandler.CreatePlayer)       //プレイヤー作成
	e.PUT("/api/players/:id", playerHandler.UpdatePlayer)    //プレイヤー情報編集更新
	e.DELETE("/api/players/:id", playerHandler.DeletePlayer) //プレイヤー削除
	e.GET("/players/:id/calculate_elo", playerHandler.CalculateEloRatingHandler)
	e.POST("/players", playerHandler.CreatePlayerHandler) //プレイヤー作成
}
