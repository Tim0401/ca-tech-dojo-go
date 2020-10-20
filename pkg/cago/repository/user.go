package repository

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"context"
	"database/sql"
)

// UserRepository ユーザーリポジトリ
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
}

type dbUserRepository struct {
	db sql.DB
}

// NewDbUserRepository ユーザーリポジトリの作成
func NewDbUserRepository(db *sql.DB) UserRepository {
	return &dbUserRepository{*db}
}

func (ur *dbUserRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Query("INSERT INTO テーブル名（列名1,列名2,……）")
}
