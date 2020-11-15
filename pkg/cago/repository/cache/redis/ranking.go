package redis

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"golang.org/x/xerrors"
)

type redisRankingRepository struct {
	conn redis.Conn
}

func (r *redisRankingRepository) Top(limit int) ([]model.UserRanking, error) {
	var userRanks []model.UserRanking
	strs, err := redis.Strings(r.conn.Do("ZREVRANGE", "user_ranking", 0, limit, "WITHSCORES"))
	if err != nil {
		return userRanks, xerrors.Errorf("Call Top: %w", err)
	}
	for i := 0; i < len(strs); i += 2 {
		var userRank model.UserRanking
		userRank.UserID, err = strconv.Atoi(strs[i])
		if err != nil {
			return userRanks, xerrors.Errorf("Call Atoi: %w", err)
		}
		userRank.Score, err = strconv.Atoi(strs[i+1])
		if err != nil {
			return userRanks, xerrors.Errorf("Call Atoi: %w", err)
		}
		userRank.Rank = (i + 2) / 2

		userRanks = append(userRanks, userRank)
	}

	return userRanks, nil
}

func (r *redisRankingRepository) Set(userRanks []model.UserRanking) error {

	conn := r.conn

	for _, userRank := range userRanks {
		conn.Send("ZADD", "user_ranking", userRank.Score, userRank.UserID)
	}
	return nil
}
func (r *redisRankingRepository) Clear() error {
	return nil
}
