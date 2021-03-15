package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/repository/cache"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"

	"golang.org/x/xerrors"
)

// RankingService RankingService
type RankingService interface {
	GetUserRanking(ctx context.Context, ranking *input.GetUserRanking) (output.GetUserRanking, error)
	UpdateUserRanking(ctx context.Context) error
}

type rankingService struct {
	r  repository.Repository
	cr cache.Repository
}

// NewRankingService NewRankingService
func NewRankingService(r repository.Repository, cr cache.Repository) RankingService {
	return &rankingService{r, cr}
}

func (rs *rankingService) GetUserRanking(ctx context.Context, ranking *input.GetUserRanking) (output.GetUserRanking, error) {
	var outputGetUserRanking output.GetUserRanking

	conn, err := rs.cr.NewConnection()
	if err != nil {
		return outputGetUserRanking, xerrors.Errorf("Call NewConnection: %w", err)
	}
	defer conn.Close()

	userRanks, err := conn.Ranking().Top(ranking.Top, ranking.Skip)
	if err != nil {
		return outputGetUserRanking, xerrors.Errorf("Call Set: %w", err)
	}

	for _, userRank := range userRanks {
		var outputUserRank output.UserRank
		outputUserRank.Rank = userRank.Rank
		outputUserRank.Score = userRank.Score
		outputUserRank.UserID = userRank.UserID
		outputGetUserRanking.Ranks = append(outputGetUserRanking.Ranks, outputUserRank)
	}

	return outputGetUserRanking, nil
}

func (rs *rankingService) UpdateUserRanking(ctx context.Context) error {

	db, err := rs.r.NewConnection()
	if err != nil {
		return xerrors.Errorf("Call NewConnection: %w", err)
	}

	userRanks, err := db.User().GetAllUserScore()
	if err != nil {
		return xerrors.Errorf("Call GetAllUserScore: %w", err)
	}

	conn, err := rs.cr.NewConnection()
	if err != nil {
		return xerrors.Errorf("Call NewConnection: %w", err)
	}
	defer conn.Close()

	err = conn.RunTransaction(func(tx cache.Transaction) error {
		err := tx.Ranking().Set(userRanks)
		if err != nil {
			return xerrors.Errorf("Call Set: %w", err)
		}

		return nil
	})

	if err != nil {
		return xerrors.Errorf("Call RunTransaction: %w", err)
	}

	return nil
}
