package presenter

import (
	"ca-tech-dojo-go/pkg/cago/presenter/input"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CharaPresenter CharaPresenter
type RankingPresenter interface {
	GetUserRanking(ctx context.Context, ranking *input.GetUserRanking, w http.ResponseWriter)
	UpdateUserRanking(ctx context.Context, w http.ResponseWriter)
	ShowError(ctx context.Context, err *input.ShowError, w http.ResponseWriter)
}

type rankingPresenter struct {
}

// NewCharaPresenter NewCharaPresenter
func NewRankingPresenter() RankingPresenter {
	return &rankingPresenter{}
}

func (rp *rankingPresenter) GetUserRanking(ctx context.Context, ranking *input.GetUserRanking, w http.ResponseWriter) {
	res, err := json.Marshal(ranking)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rp *rankingPresenter) UpdateUserRanking(ctx context.Context, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (rp *rankingPresenter) ShowError(ctx context.Context, err *input.ShowError, w http.ResponseWriter) {
	// output
	if err.E != nil {
		fmt.Printf("%+v\n", err.E)
		http.Error(w, err.E.Error(), err.Status)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
