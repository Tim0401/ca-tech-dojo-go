package controller

import (
	cInput "ca-tech-dojo-go/pkg/cago/controller/input"
	"ca-tech-dojo-go/pkg/cago/interactor"
	iInput "ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type GachaController interface {
	DrawGacha(w http.ResponseWriter, r *http.Request)
}

type gachaController struct {
	gi interactor.GachaInteractor
}

// NewGachaController Gachaコントローラー作成
func NewGachaController(gi interactor.GachaInteractor) GachaController {
	return &gachaController{gi}
}

func (gc *gachaController) DrawGacha(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var controllerInput cInput.DrawGacha
	if err := json.NewDecoder(r.Body).Decode(&controllerInput); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("stub drawGacha")

	ctx := r.Context()
	modelUser, ok := ctx.Value(model.UserKey).(model.User)
	if !ok {
		fmt.Printf("ctxから取得した値をmodel.Userに変換できません。")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var interactorInput iInput.DrawGacha
	interactorInput.Times = controllerInput.Times
	interactorInput.UserID = modelUser.ID
	gc.gi.DrawGacha(ctx, &interactorInput)

	// var interactorUser iInput.UpdateUser
	// interactorUser.ID = modelUser.ID
	// interactorUser.Name = controllerUser.Name
	// // todo エラー
	// if _, err := uc.ui.UpdateUser(ctx, &interactorUser); err != nil {
	// 	var presenterError pInput.ShowError
	// 	presenterError.E = err
	// 	uc.up.ShowError(ctx, &presenterError, w)
	// }

	// var presenterUser pInput.UpdateUser
	// uc.up.UpdateUser(ctx, &presenterUser, w)
}
