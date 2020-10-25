package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/presenter"
	"ca-tech-dojo-go/pkg/cago/repository/database"
	"ca-tech-dojo-go/pkg/cago/service"
	"ca-tech-dojo-go/pkg/util"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Router ルーティング
func Router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "hello world.")
		return
	}

	config, err := util.LoadConfigForYaml()
	if err != nil {
		panic(err.Error())
	}
	db, err := sql.Open(config.Db.DbDriver, config.Db.Dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ur := repository.NewDbUserRepository(db)
	ur := database.NewRepository(db)
	us := service.NewUserService(ur)
	up := presenter.NewUserPresenter()
	ui := interactor.NewUserInteractor(us, up)
	uc := controller.NewUserController(ui)

	switch r.URL.Path {
	case "/user/create":
		if r.Method == http.MethodPost {
			uc.CreateUser(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	case "/user/get":
		if r.Method == http.MethodGet {
			uc.GetUser(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found\n")
	}
}
