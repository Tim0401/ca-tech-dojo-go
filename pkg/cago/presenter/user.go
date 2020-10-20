package presenter

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"context"
)

// UserPresenter ユーザープレゼンター
type UserPresenter interface {
	Output(ctx context.Context, user *model.User)
}

type userPresenter struct {
}

// NewUserService ユーザーサービス作成
func NewUserService() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) Output(ctx context.Context, user *model.User) {
	// output
}
