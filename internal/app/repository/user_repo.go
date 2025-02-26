package repository

import "fmt"

// UserRepository はデータアクセスのインターフェース
type UserRepository interface {
    GetUser(id int) string
}

// userRepo は UserRepository の実装
type userRepo struct{}

// NewUserRepository は UserRepository のインスタンスを作成
func NewUserRepository() UserRepository {
    return &userRepo{}
}

// GetUser は ID に基づいてユーザー名を取得する
func (r *userRepo) GetUser(id int) string {
    return fmt.Sprintf("User%d", id)
}
