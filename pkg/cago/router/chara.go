package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

type CharaRouter interface {
	CharaRouter(w http.ResponseWriter, r *http.Request)
}

type charaRouter struct {
	cc controller.CharaController
}

// NewCharaRouter NewCharaRouter
func NewCharaRouter(cc controller.CharaController) CharaRouter {
	return &charaRouter{cc}
}

// UserRouter ルーティング
func (cr *charaRouter) CharaRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/character/list":
		if r.Method == http.MethodGet {
			cr.cc.GetCharaList(w, r)
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
