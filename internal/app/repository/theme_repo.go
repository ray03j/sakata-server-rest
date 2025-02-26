package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// テーマのデータを保持する構造体
type Theme struct {
	ID    string `json:"id"`
	Theme string `json:"theme"`
}

// テーマを管理するリポジトリ
type ThemeRepository struct {
	mu     sync.RWMutex
	themes map[string]Theme
}

// 新しいリポジトリを作成
func NewThemeRepository() *ThemeRepository {
	return &ThemeRepository{
		themes: make(map[string]Theme),
	}
}

// テーマを追加して UUID を返す
func (r *ThemeRepository) CreateTheme(themeText string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()
	r.themes[id] = Theme{ID: id, Theme: themeText}
	return id
}

// UUID に対応するテーマを取得
func (r *ThemeRepository) GetTheme(id string) (Theme, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	theme, exists := r.themes[id]
	if !exists {
		return Theme{}, errors.New("theme not found")
	}
	return theme, nil
}
