package presenter

import (
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"encoding/json"
	"net/http"
)

// UserPresenter ユーザープレゼンター
type UserPresenter interface {
	CreateUser(ctx context.Context, user *output.CreateUser, w http.ResponseWriter)
}

type userPresenter struct {
}

// NewUserPresenter ユーザーサービス作成
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) CreateUser(ctx context.Context, user *output.CreateUser, w http.ResponseWriter) {
	// output
	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

	w.WriteHeader(http.StatusCreated)
}
