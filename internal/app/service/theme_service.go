package service

import "sakata-server-rest/internal/app/repository"

// テーマのサービス
type ThemeService struct {
	repo *repository.ThemeRepository
}

// 新しいサービスを作成
func NewThemeService(repo *repository.ThemeRepository) *ThemeService {
	return &ThemeService{repo: repo}
}

// テーマを作成して UUID を返す
func (s *ThemeService) CreateTheme(themeText string) string {
	return s.repo.CreateTheme(themeText)
}

// UUID からテーマを取得
func (s *ThemeService) GetTheme(id string) (repository.Theme, error) {
	return s.repo.GetTheme(id)
}
