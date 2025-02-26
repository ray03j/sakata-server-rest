package service

import "sakata-server-rest/internal/app/repository"

// UserService はビジネスロジックを定義
type UserService interface {
    GetUserName(id int) string
}

// userService は UserService の実装
type userService struct {
    repo repository.UserRepository
}

// NewUserService は UserService のインスタンスを作成
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

// GetUserName はリポジトリを使ってユーザー名を取得
func (s *userService) GetUserName(id int) string {
    return s.repo.GetUser(id)
}
