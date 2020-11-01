package interactor

import (
	"ca-tech-dojo-go/pkg/cago/interactor/input"
	"ca-tech-dojo-go/pkg/cago/interactor/output"
	"ca-tech-dojo-go/pkg/cago/service"
	sInput "ca-tech-dojo-go/pkg/cago/service/input"
	"context"
)

// GachaInteractor GachaInteractor
type GachaInteractor interface {
	DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error)
}

type gachaInteractor struct {
	gs service.GachaService
	cs service.CharaService
}

// NewGachaInteractor NewGachaInteractor
func NewGachaInteractor(gs service.GachaService, cs service.CharaService) GachaInteractor {
	return &gachaInteractor{gs, cs}
}

func (gi *gachaInteractor) DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error) {

	var getRateInput sInput.GetGachaRate
	var output output.DrawGacha

	// ガチャレート取得
	getRateOutput, err := gi.gs.GetGachaRate(ctx, &getRateInput)
	if err != nil {
		return output, err
	}

	// ガチャ
	var drawGachaInput sInput.DrawGacha
	var charaIDs []int32
	drawGachaInput.CharaRates = getRateOutput.CharaRates
	for i := 0; i < int(gacha.Times); i++ {
		drawGachaOutput, err := gi.gs.DrawGacha(ctx, &drawGachaInput)
		if err != nil {
			return output, err
		}
		charaIDs = append(charaIDs, drawGachaOutput.CharaID)
	}

	// 登録

	// キャラクター名取得
	var getCharasInput sInput.GetCharas
	getCharasInput.IDs = charaIDs
	gi.cs.GetCharas(ctx, &getCharasInput)

	return output, nil
}
