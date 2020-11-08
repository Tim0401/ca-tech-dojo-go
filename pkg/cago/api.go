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

	_ "github.com/go-sql-driver/mysql"
)

// Serve エントリーポイント
func Serve(config *util.Config) {

	db, err := sql.Open(config.Db.DbDriver, config.Db.Dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	repository := database.NewRepository(db)

	// user
	us := service.NewUserService(repository)
	ui := interactor.NewUserInteractor(us)
	up := presenter.NewUserPresenter()
	uc := controller.NewUserController(ui, up)
	ur := router.NewUserRouter(uc)

	// gacha
	gs := service.NewGachaService(repository)
	cs := service.NewCharaService(repository)
	gi := interactor.NewGachaInteractor(gs, cs)
	gp := presenter.NewGachaPresenter()
	gc := controller.NewGachaController(gi, gp)
	gr := router.NewGachaRouter(gc)

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(repository)
	middlewares := middleware.NewMws(authMiddleware)

	mux := http.NewServeMux()

	mux.HandleFunc("/user/create", ur.UserRouter)
	mux.HandleFunc("/user/", middlewares.Then(ur.UserRouter))
	mux.HandleFunc("/gacha/", middlewares.Then(gr.GachaRouter))
	http.ListenAndServe(":8080", mux)
}
