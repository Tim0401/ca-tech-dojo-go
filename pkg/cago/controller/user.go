package controller

import (
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/presenter"
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
	up presenter.UserPresenter
}

// NewUserController ユーザーコントローラー作成
func NewUserController(ui interactor.UserInteractor, up presenter.UserPresenter) UserController {
	return &userController{ui, up}
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
	// todo エラー
	output, _ := uc.ui.CreateUser(ctx, &user)
	uc.up.CreateUser(ctx, &output, w)
}

func (uc *userController) GetUser(w http.ResponseWriter, r *http.Request) {

	var user input.GetUser
	token := r.Header.Get("x-token")

	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Xtoken = token
	ctx := r.Context()
	// todo エラー
	output, _ := uc.ui.GetUser(ctx, &user)
	uc.up.GetUser(ctx, &output, w)
}
func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user input.UpdateUser

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
	err = json.Unmarshal(body[:length], &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("x-token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Xtoken = token

	ctx := r.Context()
	// todo エラー
	output, _ := uc.ui.UpdateUser(ctx, &user)
	uc.up.UpdateUser(ctx, &output, w)
}
