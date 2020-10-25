package repository

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"context"
	"database/sql"
)

// UserRepository ユーザーリポジトリ
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
}

type dbUserRepository struct {
	db sql.DB
}

// NewDbUserRepository ユーザーリポジトリの作成
func NewDbUserRepository(db *sql.DB) UserRepository {
	return &dbUserRepository{*db}
}

func (ur *dbUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	sql := "INSERT INTO user (name, token, created_at) VALUES (?, ?, ?)"
	_, err := ur.db.Exec(sql, user.Name, user.Token, user.CreatedAt)
	return err
}
