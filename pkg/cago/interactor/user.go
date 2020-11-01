package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
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
	var serviceInput sInput.GetUser
	var output output.GetUser
	serviceInput.ID = user.ID
	serviceOutput, _ := ui.us.GetUser(ctx, &serviceInput)
	output.Name = serviceOutput.Name
	return output, nil
}

func (ui *userInteractor) CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error) {
	// todo errチェック
	var serviceInput sInput.CreateUser
	var output output.CreateUser
	serviceInput.Name = user.Name
	serviceOutput, _ := ui.us.CreateUser(ctx, &serviceInput)
	output.Xtoken = serviceOutput.Xtoken
	return output, nil
}

func (ui *userInteractor) UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error) {
	// todo errチェック
	var serviceInput sInput.UpdateUser
	var output output.UpdateUser
	serviceInput.ID = user.ID
	serviceInput.Name = user.Name
	if _, err := ui.us.UpdateUser(ctx, &serviceInput); err != nil {
		return output, err
	}
	return output, nil
}
