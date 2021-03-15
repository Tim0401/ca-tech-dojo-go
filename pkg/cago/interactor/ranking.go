package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
	sOutput "ca-tech-dojo-go/pkg/cago/service/output"
	"context"

	"golang.org/x/xerrors"
)

// RankingInteractor RankingInteractor
type RankingInteractor interface {
	GetUserRanking(ctx context.Context, ranking *input.GetUserRanking) (output.GetUserRanking, error)
	UpdateUserRanking(ctx context.Context) error
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
	sOutputGetUserRanking, err := ri.rs.GetUserRanking(ctx, &sInput.GetUserRanking{Top: ranking.Top, Skip: ranking.Skip})
	if err != nil {
		return outputGetUserRanking, xerrors.Errorf("Call GetUserRanking: %w", err)
	}

	var sInputs []*sInput.GetUser
	for _, ranking := range sOutputGetUserRanking.Ranks {
		var userRank output.UserRank
		userRank.Rank = ranking.Rank
		userRank.Score = ranking.Score
		userRank.UserID = ranking.UserID
		userRank.UserName = ""
		outputGetUserRanking.Ranks = append(outputGetUserRanking.Ranks, &userRank)

		sInputs = append(sInputs, &sInput.GetUser{
			ID: userRank.UserID,
		})
	}
	// ユーザー名取得
	sOutputs, err := ri.us.GetUsers(ctx, sInputs)
	if err != nil {
		return outputGetUserRanking, xerrors.Errorf("Call GetUsers: %w", err)
	}

	// ユーザー名を格納
	userIDMap := make(map[int]*sOutput.GetUser)
	for _, user := range sOutputs {
		userIDMap[user.ID] = user
	}
	for _, ranking := range outputGetUserRanking.Ranks {
		if user, ok := userIDMap[ranking.UserID]; ok {
			ranking.UserName = user.Name
		}
	}

	return outputGetUserRanking, nil
}

func (ri *rankingInteractor) UpdateUserRanking(ctx context.Context) error {
	err := ri.rs.UpdateUserRanking(ctx)
	if err != nil {
		return xerrors.Errorf("Call UpdateUserRanking: %w", err)
	}
	return nil
}
