package cago

import (
	"ca-tech-dojo-go/pkg/cago/middleware"
	"ca-tech-dojo-go/pkg/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewDB(config *util.Config) *sql.DB {
	db, err := sql.Open(config.Db.DbDriver, config.Db.Dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func NewRedis(config *util.Config) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(config.Redis.Protocol, config.Redis.Address+":"+config.Redis.Port)
		},
	}

	return pool
}

func NewConfig() *util.Config {
	// config読み込み
	config, err := util.LoadConfigForYaml()
	if err != nil {
		panic(err.Error())
	}
	return config
}

// Serve エントリーポイント
func Serve() {

	ur := InitUserRouter()
	gr := InitGachaRouter()
	cr := InitCharaRouter()
	rr := InitRankingRouter()

	authMiddleware := InitMiddleware()
	middlewares := middleware.NewMws(authMiddleware)

	mux := http.NewServeMux()

	mux.HandleFunc("/user/create", ur.UserRouter)
	mux.HandleFunc("/user/", middlewares.Then(ur.UserRouter))
	mux.HandleFunc("/gacha/", middlewares.Then(gr.GachaRouter))
	mux.HandleFunc("/character/", middlewares.Then(cr.CharaRouter))
	mux.HandleFunc("/ranking/", middlewares.Then(rr.RankingRouter))

	mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)
	http.ListenAndServe(":8080", mux)
}
