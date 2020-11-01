package router

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	case "/user/get":
		if r.Method == http.MethodGet {
			ur.uc.GetUser(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	case "/user/update":
		if r.Method == http.MethodPut {
			ur.uc.UpdateUser(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found\n")
	}
}
