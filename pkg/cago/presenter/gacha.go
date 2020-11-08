package presenter

import (
	"ca-tech-dojo-go/pkg/cago/presenter/input"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GachaPresenter GachaPresenter
type GachaPresenter interface {
	DrawGacha(ctx context.Context, gacha *input.DrawGacha, w http.ResponseWriter)
	ShowError(ctx context.Context, gacha *input.ShowError, w http.ResponseWriter)
}

type gachaPresenter struct {
}

// NewGachaPresenter NewGachaPresenter
func NewGachaPresenter() GachaPresenter {
	return &gachaPresenter{}
}

func (gp *gachaPresenter) DrawGacha(ctx context.Context, gacha *input.DrawGacha, w http.ResponseWriter) {
	res, err := json.Marshal(gacha)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gp *gachaPresenter) ShowError(ctx context.Context, err *input.ShowError, w http.ResponseWriter) {
	// output
	if err.E != nil {
		fmt.Printf("%+v\n", err.E)
		http.Error(w, err.E.Error(), err.Status)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
