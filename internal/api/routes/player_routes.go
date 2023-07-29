package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/api/handler"
	"github.com/labstack/echo/v4"
)

func RegisterPlayerRoutes(e *echo.Echo, playerHandler handlers.PlayerHandler) {
	// e.GET("/api/players", playerHandler.GetAllPlayers)       //ユーザー全取得
	// e.GET("/api/players/:id", playerHandler.GetPlayerByID)   //ユーザー個別取得
	e.POST("/api/players", playerHandler.CreatePlayer) //ユーザー作成
	// e.PUT("/api/players/:id", playerHandler.UpdatePlayer)    //ユーザー情報編集更新
	// e.DELETE("/api/players/:id", playerHandler.DeletePlayer) //ユーザー削除
}
