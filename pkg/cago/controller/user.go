package controller

import (
	cInput "ca-tech-dojo-go/pkg/cago/controller/input"
	"ca-tech-dojo-go/pkg/cago/interactor"
	iInput "ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/presenter"
	pInput "ca-tech-dojo-go/pkg/cago/presenter/input"
	"encoding/json"
	"fmt"
	"net/http"
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

	//parse json
	var controllerUser cInput.CreateUser
	var presenterUser pInput.CreateUser
	var interactorUser iInput.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&controllerUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	interactorUser.Name = controllerUser.Name
	// todo エラー
	output, _ := uc.ui.CreateUser(ctx, &interactorUser)
	presenterUser.Xtoken = output.Xtoken
	uc.up.CreateUser(ctx, &presenterUser, w)
}

func (uc *userController) GetUser(w http.ResponseWriter, r *http.Request) {

	var presenterUser pInput.GetUser
	var interactorUser iInput.GetUser
	ctx := r.Context()
	modelUser, ok := ctx.Value(model.UserKey).(model.User)
	if !ok {
		fmt.Printf("ctxから取得した値をmodel.Userに変換できません。")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	interactorUser.ID = modelUser.ID
	output, _ := uc.ui.GetUser(ctx, &interactorUser)
	presenterUser.Name = output.Name
	uc.up.GetUser(ctx, &presenterUser, w)
}
func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//parse json
	var controllerUser cInput.UpdateUser
	var presenterUser pInput.UpdateUser
	var interactorUser iInput.UpdateUser
	if err := json.NewDecoder(r.Body).Decode(&controllerUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	modelUser, ok := ctx.Value(model.UserKey).(model.User)
	if !ok {
		fmt.Printf("ctxから取得した値をmodel.Userに変換できません。")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	interactorUser.ID = modelUser.ID
	interactorUser.Name = controllerUser.Name
	// todo エラー
	if _, err := uc.ui.UpdateUser(ctx, &interactorUser); err != nil {

	}
	uc.up.UpdateUser(ctx, &presenterUser, w)
}
