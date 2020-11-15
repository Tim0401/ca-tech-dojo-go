package controller

import (
	"ca-tech-dojo-go/pkg/cago/interactor"
	iInput "ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/presenter"
	pInput "ca-tech-dojo-go/pkg/cago/presenter/input"
	"net/http"

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

	var interactorInput iInput.GetUserRanking
	outputGetUserRanking, err := rc.ri.GetUserRanking(ctx, &interactorInput)
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
