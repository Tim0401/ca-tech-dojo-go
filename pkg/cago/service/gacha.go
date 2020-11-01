package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
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
		var charaRate output.CharaRate
		charaRate.CharaID = gachaModel.CharaID
		charaRate.Rate = gachaModel.Rate
		outputGachaRate.CharaRates = append(outputGachaRate.CharaRates, charaRate)
	}

	return outputGachaRate, err
}

func (gs *gachaService) DrawGacha(ctx context.Context, user *input.DrawGacha) (output.DrawGacha, error) {
	var outputDrawGacha output.DrawGacha

	return outputDrawGacha, nil
}
