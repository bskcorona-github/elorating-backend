// backend/handler/elo_handler.go

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

// CalculateEloRatingHandler はEloレーティングを計算するハンドラーです。
func (h *EloHandler) CalculateEloRatingHandler(c echo.Context) error {
	// TODO: Eloレーティングを計算する処理を実装する
	return c.String(http.StatusOK, "Eloレーティングを計算します")
}
