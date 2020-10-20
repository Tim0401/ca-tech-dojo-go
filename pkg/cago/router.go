package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"ca-tech-dojo-go/pkg/cago/repository"
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

	ur := repository.NewDbUserRepository(db)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)

	switch r.URL.Path {
	case "/user/create":
		if r.Method == http.MethodPost {
			uc.CreateUser(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Method not allowed.\n")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found\n")
	}
}
