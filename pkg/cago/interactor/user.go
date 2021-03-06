package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
	"context"

	"golang.org/x/xerrors"
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
	var serviceInput sInput.GetUser
	var output output.GetUser
	serviceInput.ID = user.ID
	serviceOutput, err := ui.us.GetUser(ctx, &serviceInput)
	if err != nil {
		return output, xerrors.Errorf("Call GetUser: %w", err)
	}
	output.Name = serviceOutput.Name
	return output, nil
}

func (ui *userInteractor) CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error) {
	var serviceInput sInput.CreateUser
	var output output.CreateUser
	serviceInput.Name = user.Name
	serviceOutput, err := ui.us.CreateUser(ctx, &serviceInput)
	if err != nil {
		return output, xerrors.Errorf("Call CreateUser: %w", err)
	}
	output.Xtoken = serviceOutput.Xtoken
	return output, nil
}

func (ui *userInteractor) UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error) {
	var serviceInput sInput.UpdateUser
	var output output.UpdateUser
	serviceInput.ID = user.ID
	serviceInput.Name = user.Name
	if _, err := ui.us.UpdateUser(ctx, &serviceInput); err != nil {
		return output, xerrors.Errorf("Call UpdateUser: %w", err)
	}
	return output, nil
}
