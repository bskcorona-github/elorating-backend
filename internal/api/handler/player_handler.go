package handlers

import (
	"net/http"

	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/models"
	"github.com/bskcorona-github/EloRatingSystem5vs5/elorating-backend/backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type PlayerHandler struct {
	playerUsecase usecase.PlayerUsecase
}

func NewPlayerHandler(playerUsecase usecase.PlayerUsecase) *PlayerHandler {
	return &PlayerHandler{
		playerUsecase: playerUsecase,
	}
}

func (h *PlayerHandler) CreatePlayer(c echo.Context) error {
	player := new(models.Player)
	if err := c.Bind(player); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	createdPlayer, err := h.playerUsecase.CreatePlayer(player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create player"})
	}

	return c.JSON(http.StatusCreated, createdPlayer)
}
