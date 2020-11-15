package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
	"context"

	"golang.org/x/xerrors"
)

// RankingInteractor RankingInteractor
type RankingInteractor interface {
	GetUserRanking(ctx context.Context, ranking *input.GetUserRanking) (output.GetUserRanking, error)
}

type rankingInteractor struct {
	rs service.RankingService
	us service.UserService
}

// NewRankingInteractor NewRankingInteractor
func NewRankingInteractor(rs service.RankingService, us service.UserService) RankingInteractor {
	return &rankingInteractor{rs, us}
}

func (ri *rankingInteractor) GetUserRanking(ctx context.Context, ranking *input.GetUserRanking) (output.GetUserRanking, error) {
	var outputGetUserRanking output.GetUserRanking

	// ユーザーランキングを取得
	var sInputGetUserRanking sInput.GetUserRanking
	sOutputGetUserRanking, err := ri.rs.GetUserRanking(ctx, &sInputGetUserRanking)
	if err != nil {
		return outputGetUserRanking, xerrors.Errorf("Call GetUserRanking: %w", err)
	}

	for _, ranking := range sOutputGetUserRanking.Ranks {
		var userRank output.UserRank
		userRank.Rank = ranking.Rank
		userRank.Score = ranking.Score
		userRank.UserID = ranking.UserID
		userRank.UserName = ""
		outputGetUserRanking.Ranks = append(outputGetUserRanking.Ranks, userRank)
	}
	// ユーザー名取得

	return outputGetUserRanking, nil
}
