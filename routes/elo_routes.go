// backend/routes/elo_routes.go

package routes

import (
	handlers "github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/handler"
	"github.com/labstack/echo/v4"
)

// RegisterEloRoutes はElo関連のルーティングを登録します。
func RegisterEloRoutes(e *echo.Echo, eloHandler handlers.EloHandler) {
	e.GET("/elo/calculate", eloHandler.CalculateEloRatingHandler)
	// 他のElo関連のルートをここに追加する
}
