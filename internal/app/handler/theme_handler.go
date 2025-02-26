package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakata-server-rest/internal/app/service"
)

// JSON のリクエスト/レスポンス用の構造体
type CreateThemeRequest struct {
	Theme string `json:"theme"`
}

type CreateThemeResponse struct {
	ID string `json:"id"`
}

type GetThemeResponse struct {
	ID    string `json:"id"`
	Theme string `json:"theme"`
}

// テーマのハンドラー
type ThemeHandler struct {
	service *service.ThemeService
}

// 新しいハンドラーを作成
func NewThemeHandler(service *service.ThemeService) *ThemeHandler {
	return &ThemeHandler{service: service}
}

// テーマを作成
func (h *ThemeHandler) CreateTheme(c *gin.Context) {
	var req CreateThemeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := h.service.CreateTheme(req.Theme)
	c.JSON(http.StatusOK, CreateThemeResponse{ID: id})
}

// テーマを取得
func (h *ThemeHandler) GetTheme(c *gin.Context) {
	id := c.Param("id")

	theme, err := h.service.GetTheme(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Theme not found"})
		return
	}

	c.JSON(http.StatusOK, GetThemeResponse{ID: theme.ID, Theme: theme.Theme})
}
