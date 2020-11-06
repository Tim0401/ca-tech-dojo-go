package service

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/io"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"

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

// NewGachaService Gachaサービス作成
func NewGachaService(r repository.Repository) GachaService {
	return &gachaService{r}
}

func (gs *gachaService) GetGachaRate(ctx context.Context, gacha *input.GetGachaRate) (output.GetGachaRate, error) {
	var outputGachaRate output.GetGachaRate

	con, err := gs.r.NewConnection()
	if err != nil {
		return outputGachaRate, err
	}
	defer con.Close()

	gachaModels, err := con.Gacha().FindByGachaType(gacha.GachaType)
	if err != nil {
		return outputGachaRate, err
	}

	// 格納
	charaRates := make(map[int]*io.CharaRates)
	for _, gachaModel := range gachaModels {
		var charaRate io.CharaRate
		charaRate.CharaID = gachaModel.CharaID
		charaRate.RateTypeID = gachaModel.RateTypeID
		charaRate.Rate = gachaModel.Rate
		if _, ok := charaRates[int(charaRate.RateTypeID)]; !ok {
			charaRates[int(charaRate.RateTypeID)] = new(io.CharaRates)
		}
		charaRates[int(charaRate.RateTypeID)].SumRate += charaRate.Rate
		charaRates[int(charaRate.RateTypeID)].CharaRateArray = append(charaRates[int(charaRate.RateTypeID)].CharaRateArray, &charaRate)
	}
	outputGachaRate.CharaRates = charaRates

	rateTypeModels, err := con.RateType().FindByGachaType(gacha.GachaType)
	if err != nil {
		return outputGachaRate, err
	}

	// 格納
	outputGachaRate.RateTypes = new(io.RateTypes)
	for _, rateTypeModel := range rateTypeModels {
		var rateType io.RateType
		rateType.ID = rateTypeModel.ID
		rateType.Rate = rateTypeModel.Rate
		outputGachaRate.RateTypes.SumRate += rateType.Rate
		outputGachaRate.RateTypes.RateTypeArray = append(outputGachaRate.RateTypes.RateTypeArray, &rateType)
	}

	return outputGachaRate, err
}

func (gs *gachaService) DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error) {
	var outputDrawGacha output.DrawGacha

	// どのレートタイプに当たるか抽選
	sum := 0
	targetNum := rand.Intn(gacha.RateTypes.SumRate)
	var winningRateType int
	for _, rateType := range gacha.RateTypes.RateTypeArray {
		sum += rateType.Rate
		if sum > targetNum {
			winningRateType = rateType.ID
			break
		}
	}

	// レートの中のどのキャラに当たるか抽選
	sum = 0
	targetNum = rand.Intn(gacha.CharaRates[winningRateType].SumRate)
	for _, charaRate := range gacha.CharaRates[winningRateType].CharaRateArray {
		sum += charaRate.Rate
		if sum > targetNum {
			outputDrawGacha.CharaID = charaRate.CharaID
			break
		}
	}

	return outputDrawGacha, nil
}
