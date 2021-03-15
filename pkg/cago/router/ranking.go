package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

type RankingRouter interface {
	RankingRouter(w http.ResponseWriter, r *http.Request)
}

type rankingRouter struct {
	rc controller.RankingController
}

// NewRankingRouter NewRankingRouter
func NewRankingRouter(rc controller.RankingController) RankingRouter {
	return &rankingRouter{rc}
}

// rankingRouter ルーティング
func (rr *rankingRouter) RankingRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ranking/user":
		if r.Method == http.MethodGet {
			rr.rc.GetUserRanking(w, r)
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
