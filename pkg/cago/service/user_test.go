package service

import (
	"ca-tech-dojo-go/pkg/cago/repository/mock_repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetUser(t *testing.T) {
	// コントローラーの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_repository.NewMockRepository(ctrl)

	r.EXPECT().NewConnection().Return(mock_repository.MockConnection{}, nil)
	us := NewUserService(r)

	var input input.GetUser
	us.GetUser(context.Background(), &input)
}
