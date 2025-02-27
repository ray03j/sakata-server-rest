//go:build wireinject

package di

import (
    "github.com/google/wire"
    "sakata-server-rest/internal/app/handler"
    "sakata-server-rest/internal/app/repository"
    "sakata-server-rest/internal/app/service"
)

// InitializeUserHandler は DI を適用した UserHandler を作成
func InitializeUserHandler() *handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
	)
    return nil // 実際には Wire がこの値を無視して適切な依存関係を解決する
}

// InitializeThemeHandler は DI で ThemeHandler を生成
func InitializeThemeHandler() *handler.ThemeHandler {
	wire.Build(
		repository.NewThemeRepository,
		service.NewThemeService,
		handler.NewThemeHandler,
	)
	return &handler.ThemeHandler{}
}

func InitializeGameHandler() *handler.GameHandler {
	wire.Build(
		repository.NewDB,
		repository.NewGameRepository,
		service.NewGameService,
		handler.NewGameHandler,
	)
	return nil
}
