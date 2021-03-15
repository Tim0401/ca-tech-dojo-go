package service

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository/mock_repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// コントローラーの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := 1

	r := mock_repository.NewMockRepository(ctrl)
	c := mock_repository.NewMockConnection(ctrl)
	uq := mock_repository.NewMockUserQuery(ctrl)

	r.EXPECT().NewConnection().Return(c, nil)
	c.EXPECT().User().Return(uq)
	uq.EXPECT().Find(id).Return(model.User{
		ID:        1,
		Name:      "testuser1",
		Token:     "testtoken1",
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
	}, nil)

	us := NewUserService(r)

	user, err := us.GetUser(context.Background(), &input.GetUser{ID: id})
	if err != nil {
		t.Fatal("err GetUser", err)
	}

	assert.Exactly(t, user, output.GetUser{
		Name: "testuser1",
	})
}
