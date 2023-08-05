package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

func RegisterPlayerRoutes(e *echo.Echo, h *handlers.PlayerHandler) {
	e.GET("/api/players", h.GetAllPlayers)       //プレイヤー全取得
	e.GET("/api/players/:id", h.GetPlayerByID)   //プレイヤー個別取得
	e.POST("/api/players", h.CreatePlayer)       //プレイヤー作成
	e.PUT("/api/players/:id", h.UpdatePlayer)    //プレイヤー情報編集更新
	e.DELETE("/api/players/:id", h.DeletePlayer) //プレイヤー削除

}
