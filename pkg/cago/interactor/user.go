package interactor

import (
	"ca-tech-dojo-go/pkg/cago/presenter"
	"ca-tech-dojo-go/pkg/cago/service"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"context"
	"net/http"
)

// UserInteractor ユーザーサービス
type UserInteractor interface {
	CreateUser(ctx context.Context, user *input.CreateUser, w http.ResponseWriter)
	GetUser(ctx context.Context, user *input.GetUser, w http.ResponseWriter)
	UpdateUser(ctx context.Context, user *input.UpdateUser, w http.ResponseWriter)
}

type userInteractor struct {
	us service.UserService
	up presenter.UserPresenter
}

// NewUserInteractor ユーザーサービス作成
func NewUserInteractor(us service.UserService, up presenter.UserPresenter) UserInteractor {
	return &userInteractor{us, up}
}

func (ui *userInteractor) GetUser(ctx context.Context, user *input.GetUser, w http.ResponseWriter) {
	outputUser, _ := ui.us.GetUser(ctx, user)
	ui.up.GetUser(ctx, &outputUser, w)
}

func (ui *userInteractor) CreateUser(ctx context.Context, user *input.CreateUser, w http.ResponseWriter) {
	// todo errチェック
	outputUser, _ := ui.us.CreateUser(ctx, user)
	ui.up.CreateUser(ctx, &outputUser, w)
}

func (ui *userInteractor) UpdateUser(ctx context.Context, user *input.UpdateUser, w http.ResponseWriter) {
	// todo errチェック
	outputUser, _ := ui.us.UpdateUser(ctx, user)
	ui.up.UpdateUser(ctx, &outputUser, w)
}
