package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"
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
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found\n")
	}
}
