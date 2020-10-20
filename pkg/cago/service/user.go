package service

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository"
	"context"
)

// UserService ユーザーサービス
type UserService interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userService struct {
	ur repository.UserRepository
}

// NewUserService ユーザーサービス作成
func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) {
	us.ur.CreateUser(ctx, user)
}
