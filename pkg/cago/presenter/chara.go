package presenter

import (
	"ca-tech-dojo-go/pkg/cago/presenter/input"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CharaPresenter CharaPresenter
type CharaPresenter interface {
	GetCharaList(ctx context.Context, chara *input.GetCharaList, w http.ResponseWriter)
	ShowError(ctx context.Context, err *input.ShowError, w http.ResponseWriter)
}

type charaPresenter struct {
}

// NewCharaPresenter NewCharaPresenter
func NewCharaPresenter() CharaPresenter {
	return &charaPresenter{}
}

func (cp *charaPresenter) GetCharaList(ctx context.Context, chara *input.GetCharaList, w http.ResponseWriter) {
	res, err := json.Marshal(chara)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (cp *charaPresenter) ShowError(ctx context.Context, err *input.ShowError, w http.ResponseWriter) {
	// output
	if err.E != nil {
		fmt.Printf("%+v\n", err.E)
		http.Error(w, err.E.Error(), err.Status)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
