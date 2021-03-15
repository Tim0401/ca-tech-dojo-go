package cache

import "ca-tech-dojo-go/pkg/cago/model"

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	RunTransaction(func(tx Transaction) error) error

	Ranking() RankingQuery
}

type Transaction interface {
	Ranking() RankingCommand
}

type RankingQuery interface {
	Top(limit int, offset int) ([]model.UserRanking, error)
}

type RankingCommand interface {
	RankingQuery

	Set(userRanks []model.UserRanking) error
	Clear() error
}
