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

type GachaController interface {
	DrawGacha(w http.ResponseWriter, r *http.Request)
}

type gachaController struct {
	gi interactor.GachaInteractor
	gp presenter.GachaPresenter
}

// NewGachaController Gachaコントローラー作成
func NewGachaController(gi interactor.GachaInteractor, gp presenter.GachaPresenter) GachaController {
	return &gachaController{gi, gp}
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
	outputDrawGacha, err := gc.gi.DrawGacha(ctx, &interactorInput)
	if err != nil {
		var presenterError pInput.ShowError
		presenterError.E = err
		gc.gp.ShowError(ctx, &presenterError, w)
	}

	var presenterGacha pInput.DrawGacha
	for _, chara := range outputDrawGacha.Charas {
		var pChara pInput.Chara
		pChara.ID = chara.ID
		pChara.Name = chara.Name
		presenterGacha.Results = append(presenterGacha.Results, pChara)
	}
	gc.gp.DrawGacha(ctx, &presenterGacha, w)
}
