package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

type GachaRouter interface {
	GachaRouter(w http.ResponseWriter, r *http.Request)
}

type gachaRouter struct {
	gc controller.GachaController
}

// NewGachaRouter GachaRouter作成
func NewGachaRouter(gc controller.GachaController) GachaRouter {
	return &gachaRouter{gc}
}

// UserRouter ルーティング
func (gr *gachaRouter) GachaRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/gacha/draw":
		if r.Method == http.MethodPost {
			gr.gc.DrawGacha(w, r)
		} else {
			err := xerrors.New("Method not allowed")
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	default:
		err := xerrors.New("404 Not Found")
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}
}
