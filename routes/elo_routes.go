package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

// RegisterELORoutes はelo関連のルーティングを登録する関数です。
func RegisterELORoutes(e *echo.Echo, eloHandler *handlers.ELOHandler) {
	// プレイヤー選択API
	e.POST("/selected_players", eloHandler.SelectPlayersHandler)

	// チーム分け結果取得API
	e.GET("/teams", eloHandler.GetTeamFormationHandler)

	// ゲーム結果反映API
	e.POST("/game_result", eloHandler.ReflectGameResultHandler)
}
