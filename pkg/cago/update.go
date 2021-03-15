package cago

import "context"

// Update エントリーポイント
func Update() {
	gs := InitRankingService()
	gs.UpdateUserRanking(context.Background())
}
