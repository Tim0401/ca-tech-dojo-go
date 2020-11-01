package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/middleware"
	"ca-tech-dojo-go/pkg/cago/presenter"
	"ca-tech-dojo-go/pkg/cago/repository/database"
	"ca-tech-dojo-go/pkg/cago/router"
	"ca-tech-dojo-go/pkg/cago/service"
	"ca-tech-dojo-go/pkg/util"
	"database/sql"
	"net/http"
)

// Serve エントリーポイント
func Serve(config *util.Config) {

	db, err := sql.Open(config.Db.DbDriver, config.Db.Dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	repository := database.NewRepository(db)
	us := service.NewUserService(repository)
	ui := interactor.NewUserInteractor(us)
	up := presenter.NewUserPresenter()
	uc := controller.NewUserController(ui, up)
	ur := router.NewUserRouter(uc)

	authMiddleware := middleware.NewAuthMiddleware(repository)
	middlewares := middleware.NewMws(authMiddleware)

	mux := http.NewServeMux()

	mux.HandleFunc("/user/create", ur.UserRouter)
	mux.HandleFunc("/user/", middlewares.Then(ur.UserRouter))
	http.ListenAndServe(":8080", mux)
}
