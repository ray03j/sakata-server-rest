package handler

import (
    "net/http"
    "sakata-server-rest/internal/model"
    "sakata-server-rest/internal/app/service"

    "github.com/gin-gonic/gin"
		"github.com/google/uuid"
)

type GameHandler struct {
    service *service.GameService
}

func NewGameHandler(service *service.GameService) *GameHandler {
    return &GameHandler{service: service}
}

func (h *GameHandler) SaveGame(c *gin.Context) {
	var game model.Game
	if err := c.ShouldBindJSON(&game); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	// `id` が空の場合に `UUID` を自動生成
	if game.ID == "" {
			game.ID = uuid.New().String()
	}

	if err := h.service.SaveGame(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save game"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game saved", "gameId": game.ID})
}

func (h *GameHandler) GetGame(c *gin.Context) {
    gameID := c.Param("id")
    game, err := h.service.GetGameByID(gameID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
        return
    }

    c.JSON(http.StatusOK, game)
}

func (h *GameHandler) UpdateGame(c *gin.Context) {
	gameID := c.Param("id")
	var game model.Game
	if err := c.ShouldBindJSON(&game); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	game.ID = gameID // URLのIDを優先

	if err := h.service.UpdateGame(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game updated"})
}

