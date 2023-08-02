package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

func RegisterEloRoutes(e *echo.Echo, eloHandler *handlers.EloHandler) {
	// チーム分けAPIのエンドポイントを設定
	e.POST("/team", eloHandler.TeamFormationHandler)

	// ゲーム結果の入力APIのエンドポイントを設定
	e.POST("/result", eloHandler.GameResultHandler)
}
