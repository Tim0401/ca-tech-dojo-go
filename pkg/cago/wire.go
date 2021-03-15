//+build wireinject

package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller"
	"ca-tech-dojo-go/pkg/cago/interactor"
	"ca-tech-dojo-go/pkg/cago/middleware"
	"ca-tech-dojo-go/pkg/cago/presenter"
	"ca-tech-dojo-go/pkg/cago/repository/database"
	"ca-tech-dojo-go/pkg/cago/router"
	"ca-tech-dojo-go/pkg/cago/service"

	redisRepository "ca-tech-dojo-go/pkg/cago/repository/cache/redis"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(NewConfig, NewDB, NewRedis, database.NewRepository, redisRepository.NewRepository)

func InitUserRouter() router.UserRouter {
	wire.Build(
		SuperSet,
		service.NewUserService,
		interactor.NewUserInteractor,
		presenter.NewUserPresenter,
		controller.NewUserController,
		router.NewUserRouter,
	)
	return nil
}

func InitGachaRouter() router.GachaRouter {
	wire.Build(
		SuperSet,
		service.NewGachaService,
		service.NewCharaService,
		interactor.NewGachaInteractor,
		presenter.NewGachaPresenter,
		controller.NewGachaController,
		router.NewGachaRouter,
	)
	return nil
}

func InitRankingService() service.RankingService {
	wire.Build(
		SuperSet,
		service.NewRankingService,
	)
	return nil
}

func InitCharaRouter() router.CharaRouter {
	wire.Build(
		SuperSet,
		service.NewCharaService,
		interactor.NewCharaInteractor,
		presenter.NewCharaPresenter,
		controller.NewCharaController,
		router.NewCharaRouter,
	)
	return nil
}

func InitRankingRouter() router.RankingRouter {
	wire.Build(
		SuperSet,
		service.NewUserService,
		service.NewRankingService,
		interactor.NewRankingInteractor,
		presenter.NewRankingPresenter,
		controller.NewRankingController,
		router.NewRankingRouter,
	)
	return nil
}

func InitMiddleware() middleware.Middleware {
	wire.Build(
		SuperSet,
		middleware.NewAuthMiddleware,
	)
	return nil
}
