package presenter

import (
	"ca-tech-dojo-go/pkg/cago/presenter/input"
	"context"
	"encoding/json"
	"net/http"
)

// UserPresenter ユーザープレゼンター
type UserPresenter interface {
	CreateUser(ctx context.Context, user *input.CreateUser, w http.ResponseWriter)
	GetUser(ctx context.Context, user *input.GetUser, w http.ResponseWriter)
	UpdateUser(ctx context.Context, user *input.UpdateUser, w http.ResponseWriter)
}

type userPresenter struct {
}

// NewUserPresenter ユーザーサービス作成
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) GetUser(ctx context.Context, user *input.GetUser, w http.ResponseWriter) {
	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (up *userPresenter) CreateUser(ctx context.Context, user *input.CreateUser, w http.ResponseWriter) {
	// output
	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (up *userPresenter) UpdateUser(ctx context.Context, user *input.UpdateUser, w http.ResponseWriter) {
	// output
	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(res)
}
