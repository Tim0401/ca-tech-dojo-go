package controller

import (
	"ca-tech-dojo-go/pkg/cago/interactor"
	iInput "ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/presenter"
	pInput "ca-tech-dojo-go/pkg/cago/presenter/input"
	"net/http"

	"golang.org/x/xerrors"
)

type CharaController interface {
	GetCharaList(w http.ResponseWriter, r *http.Request)
}

type charaController struct {
	ci interactor.CharaInteractor
	cp presenter.CharaPresenter
}

// NewCharaController NewCharaController
func NewCharaController(ci interactor.CharaInteractor, cp presenter.CharaPresenter) CharaController {
	return &charaController{ci, cp}
}

func (cc *charaController) GetCharaList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	modelUser, ok := ctx.Value(model.UserKey).(model.User)
	if !ok {
		var presenterError pInput.ShowError
		presenterError.E = xerrors.New("ctxから取得した値をmodel.Userに変換できません")
		presenterError.Status = http.StatusInternalServerError
		cc.cp.ShowError(ctx, &presenterError, w)
		return
	}

	var interactorInput iInput.GetCharaList
	interactorInput.UserID = modelUser.ID
	outputGetCharaList, err := cc.ci.GetCharaList(ctx, &interactorInput)
	if err != nil {
		var presenterError pInput.ShowError
		presenterError.E = err
		presenterError.Status = http.StatusInternalServerError
		cc.cp.ShowError(ctx, &presenterError, w)
		return
	}

	var presenterChara pInput.GetCharaList

	for _, chara := range outputGetCharaList.Charas {
		var userChara pInput.UserChara
		userChara.ID = chara.ID
		userChara.CharaID = chara.CharaID
		userChara.Name = chara.Name
		presenterChara.Characters = append(presenterChara.Characters, userChara)
	}
	// 結果がない場合は空配列にする
	if presenterChara.Characters == nil {
		presenterChara.Characters = []pInput.UserChara{}
	}
	cc.cp.GetCharaList(ctx, &presenterChara, w)
}
