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

func (r *redisRankingRepository) Top(limit int, offset int) ([]model.UserRanking, error) {
	var userRanks []model.UserRanking
	strs, err := redis.Strings(r.conn.Do("ZREVRANGE", "user_ranking", offset, offset+limit-1, "WITHSCORES"))
	if err != nil {
		return userRanks, xerrors.Errorf("Call Top: %w", err)
	}

	// 結果が空の場合
	if len(strs) < 1 {
		return userRanks, nil
	}

	// ランク開始位置を取得
	score, err := strconv.ParseInt(strs[1], 10, 32)
	if err != nil {
		return userRanks, xerrors.Errorf("Call ParseInt: %w", err)
	}
	count, err := redis.Int(r.conn.Do("ZCOUNT", "user_ranking", score+1, "+inf"))
	if err != nil {
		return userRanks, xerrors.Errorf("Call ZCOUNT: %w", err)
	}
	baseRank := count + 1
	stockRank := offset

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
		// ランク判定
		stockRank += 1
		// 前のスコアと異なるならランクを現在のものに下げる
		if i > 0 && strs[i-1] != strs[i+1] {
			baseRank = stockRank
		}

		userRank.Rank = int(baseRank)
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
