package controller

import (
	"ca-tech-dojo-go/pkg/cago/interactor"
	iInput "ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/presenter"
	pInput "ca-tech-dojo-go/pkg/cago/presenter/input"
	"net/http"
	"strconv"

	"golang.org/x/xerrors"
)

type RankingController interface {
	GetUserRanking(w http.ResponseWriter, r *http.Request)
}

type rankingController struct {
	ri interactor.RankingInteractor
	rp presenter.RankingPresenter
}

// NewCharaController NewCharaController
func NewRankingController(ri interactor.RankingInteractor, rp presenter.RankingPresenter) RankingController {
	return &rankingController{ri, rp}
}

func (rc *rankingController) GetUserRanking(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	v := r.URL.Query()
	// TODO: リクエストの値のバリデーション
	top := v.Get("$Top")
	skip := v.Get("$Skip")
	intTop, err := strconv.ParseInt(top, 10, 32)
	if err != nil {
		intTop = 100
	}
	intSkip, err := strconv.ParseInt(skip, 10, 32)
	if err != nil {
		intSkip = 0
	}

	outputGetUserRanking, err := rc.ri.GetUserRanking(ctx, &iInput.GetUserRanking{Top: int(intTop), Skip: int(intSkip)})
	if err != nil {
		var presenterError pInput.ShowError
		presenterError.E = xerrors.Errorf("Call GetUserRanking: %w", err)
		presenterError.Status = http.StatusInternalServerError
		rc.rp.ShowError(ctx, &presenterError, w)
		return
	}

	var presenterInput pInput.GetUserRanking

	for _, ranking := range outputGetUserRanking.Ranks {
		var userRank pInput.UserRank
		userRank.Rank = ranking.Rank
		userRank.Score = ranking.Score
		userRank.UserID = ranking.UserID
		userRank.UserName = ranking.UserName
		presenterInput.Results = append(presenterInput.Results, userRank)
	}

	// 結果がない場合は空配列にする
	if presenterInput.Results == nil {
		presenterInput.Results = []pInput.UserRank{}
	}
	rc.rp.GetUserRanking(ctx, &presenterInput, w)
}
