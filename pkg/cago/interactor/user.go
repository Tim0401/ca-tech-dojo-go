package interactor

import (
	"ca-tech-dojo-go/pkg/cago/service"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
)

// UserInteractor ユーザーサービス
type UserInteractor interface {
	CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error)
	GetUser(ctx context.Context, user *input.GetUser) (output.GetUser, error)
	UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error)
}

type userInteractor struct {
	us service.UserService
}

// NewUserInteractor ユーザーサービス作成
func NewUserInteractor(us service.UserService) UserInteractor {
	return &userInteractor{us}
}

func (ui *userInteractor) GetUser(ctx context.Context, user *input.GetUser) (output.GetUser, error) {
	// todo errチェック
	outputUser, _ := ui.us.GetUser(ctx, user)
	return outputUser, nil
}

func (ui *userInteractor) CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error) {
	// todo errチェック
	outputUser, _ := ui.us.CreateUser(ctx, user)
	return outputUser, nil
}

func (ui *userInteractor) UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error) {
	// todo errチェック
	outputUser, _ := ui.us.UpdateUser(ctx, user)
	return outputUser, nil
}
