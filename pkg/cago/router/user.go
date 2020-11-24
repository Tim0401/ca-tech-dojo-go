package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

type UserRouter interface {
	UserRouter(w http.ResponseWriter, r *http.Request)
}

type userRouter struct {
	uc controller.UserController
}

// NewUserRouter ユーザーRouter作成
func NewUserRouter(uc controller.UserController) UserRouter {
	return &userRouter{uc}
}

// UserRouter ルーティング
func (ur *userRouter) UserRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/user/create":
		if r.Method == http.MethodPost {
			ur.uc.CreateUser(w, r)
		} else {
			err := xerrors.New("Method not allowed")
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "/user/get":
		if r.Method == http.MethodGet {
			ur.uc.GetUser(w, r)
		} else {
			err := xerrors.New("Method not allowed")
			fmt.Printf("%+v\n", err)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "/user/update":
		if r.Method == http.MethodPut {
			ur.uc.UpdateUser(w, r)
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
