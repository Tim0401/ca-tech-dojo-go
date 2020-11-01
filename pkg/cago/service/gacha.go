package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/io"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"errors"
	"math/rand"
)

// UserService ユーザーサービス
type GachaService interface {
	GetGachaRate(ctx context.Context, user *input.GetGachaRate) (output.GetGachaRate, error)
	DrawGacha(ctx context.Context, user *input.DrawGacha) (output.DrawGacha, error)
}

type gachaService struct {
	r repository.Repository
}

// NewUserService ユーザーサービス作成
func NewGachaService(r repository.Repository) GachaService {
	return &gachaService{r}
}

func (gs *gachaService) GetGachaRate(ctx context.Context, user *input.GetGachaRate) (output.GetGachaRate, error) {
	var outputGachaRate output.GetGachaRate

	con, err := gs.r.NewConnection()
	if err != nil {
		return outputGachaRate, err
	}
	defer con.Close()

	gachaModels, err := con.Gacha().FindAll()
	if err != nil {
		return outputGachaRate, err
	}

	// 格納
	for _, gachaModel := range gachaModels {
		var charaRate io.CharaRate
		charaRate.CharaID = gachaModel.CharaID
		charaRate.Rate = gachaModel.Rate
		outputGachaRate.CharaRates = append(outputGachaRate.CharaRates, charaRate)
	}

	return outputGachaRate, err
}

func (gs *gachaService) DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error) {
	var outputDrawGacha output.DrawGacha
	resultRate := rand.Float64()
	var sum = 0.0

	for _, chara := range gacha.CharaRates {
		sum += chara.Rate / 100
		if sum > resultRate {
			outputDrawGacha.CharaID = chara.CharaID
			return outputDrawGacha, nil
		}
	}

	return outputDrawGacha, errors.New("ガチャの結果が存在しません")
}
