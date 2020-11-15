package redis

import (
	"ca-tech-dojo-go/pkg/cago/repository/cache"

	"github.com/gomodule/redigo/redis"
)

// NewRepository generates a new repository using DB.
func NewRepository(pool *redis.Pool) cache.Repository {
	return &redisRepository{
		pool: pool,
	}
}

type redisRepository struct {
	pool *redis.Pool
}

type redisConnection struct {
	conn redis.Conn
}

type redisTransaction struct {
	conn redis.Conn
}

func (r *redisRepository) NewConnection() (cache.Connection, error) {
	return &redisConnection{
		conn: r.pool.Get(),
	}, nil
}

func (r *redisRepository) MustConnection() cache.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

func (con *redisConnection) Close() error {
	con.conn.Close()
	return nil
}

func (con *redisConnection) RunTransaction(f func(cache.Transaction) error) error {

	conn := con.conn
	err := conn.Send("MULTI")
	if err != nil {
		panic(err)
	}

	err = f(&redisTransaction{conn: conn})
	if err != nil {
		return err
	}

	_, err = redis.Values(conn.Do("EXEC"))
	if err != nil {
		panic(err)
	}

	return nil
}

func (con *redisConnection) Ranking() cache.RankingQuery {
	return &redisRankingRepository{conn: con.conn}
}

func (tx *redisTransaction) Ranking() cache.RankingCommand {
	return &redisRankingRepository{conn: tx.conn}
}
