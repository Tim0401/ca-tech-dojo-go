package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/middleware"
	"ca-tech-dojo-go/pkg/cago/presenter"
	redisRepository "ca-tech-dojo-go/pkg/cago/repository/cache/redis"
	"ca-tech-dojo-go/pkg/cago/repository/database"
	"ca-tech-dojo-go/pkg/cago/router"
	"ca-tech-dojo-go/pkg/cago/service"
	"ca-tech-dojo-go/pkg/util"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

// Serve エントリーポイント
func Serve(config *util.Config) {

	db, err := sql.Open(config.Db.DbDriver, config.Db.Dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(config.Redis.Protocol, config.Redis.Address+":"+config.Redis.Port)
		},
	}

	redisRepo := redisRepository.NewRepository(pool)
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

	// character
	ci := interactor.NewCharaInteractor(cs)
	cp := presenter.NewCharaPresenter()
	cc := controller.NewCharaController(ci, cp)
	cr := router.NewCharaRouter(cc)

	// ranking
	rs := service.NewRankingService(repository, redisRepo)
	ri := interactor.NewRankingInteractor(rs, us)
	rp := presenter.NewRankingPresenter()
	rc := controller.NewRankingController(ri, rp)
	rr := router.NewRankingRouter(rc)

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(repository)
	middlewares := middleware.NewMws(authMiddleware)

	mux := http.NewServeMux()

	mux.HandleFunc("/user/create", ur.UserRouter)
	mux.HandleFunc("/user/", middlewares.Then(ur.UserRouter))
	mux.HandleFunc("/gacha/", middlewares.Then(gr.GachaRouter))
	mux.HandleFunc("/character/", middlewares.Then(cr.CharaRouter))
	mux.HandleFunc("/ranking/", middlewares.Then(rr.RankingRouter))
	http.ListenAndServe(":8080", mux)
}
