package controller

import (
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	ui interactor.UserInteractor
}

// NewUserController ユーザーコントローラー作成
func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{ui}
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var user input.CreateUser
	err = json.Unmarshal(body[:length], &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	uc.ui.CreateUser(ctx, &user, w)
}

func (uc *userController) GetUser(w http.ResponseWriter, r *http.Request) {
}
func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
