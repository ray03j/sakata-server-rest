package service

import (
    "sakata-server-rest/internal/model"
    "sakata-server-rest/internal/app/repository"
)

type GameService struct {
    repo *repository.GameRepository
}

func NewGameService(repo *repository.GameRepository) *GameService {
    return &GameService{repo: repo}
}

func (s *GameService) SaveGame(game *model.Game) error {
    return s.repo.SaveGame(game)
}

func (s *GameService) GetGameByID(gameID string) (*model.Game, error) {
    return s.repo.GetGameByID(gameID)
}

func (s *GameService) UpdateGame(game *model.Game) error {
	return s.repo.UpdateGame(game)
}
