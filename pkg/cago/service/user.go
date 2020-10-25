package service

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"time"

	"github.com/google/uuid"
)

// UserService ユーザーサービス
type UserService interface {
	CreateUser(ctx context.Context, user *input.CreateUser) output.CreateUser
}

type userService struct {
	ur repository.UserRepository
}

// NewUserService ユーザーサービス作成
func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) CreateUser(ctx context.Context, user *input.CreateUser) output.CreateUser {
	var modelUser model.User
	uuidV4 := uuid.New()
	modelUser.Name = user.Name
	modelUser.Token = uuidV4.String()
	modelUser.CreatedAt = time.Now()

	us.ur.CreateUser(ctx, &modelUser)
	// エラーをどう処理するか

	var outputUser output.CreateUser
	outputUser.Xtoken = modelUser.Token
	return outputUser
}
