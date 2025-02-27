package repository

import (
    "sakata-server-rest/internal/model"
    "gorm.io/gorm"
)

type GameRepository struct {
    db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
    return &GameRepository{db: db}
}

func (r *GameRepository) SaveGame(game *model.Game) error {
    return r.db.Save(game).Error
}

func (r *GameRepository) GetGameByID(gameID string) (*model.Game, error) {
    var game model.Game
    if err := r.db.First(&game, "id = ?", gameID).Error; err != nil {
        return nil, err
    }
    return &game, nil
}

func (r *GameRepository) UpdateGame(game *model.Game) error {
	return r.db.Model(&model.Game{}).
			Where("id = ?", game.ID).
			Updates(map[string]interface{}{
					"human_score": game.HumanScore,
					"ai_score":    game.AIScore,
			}).Error
}
