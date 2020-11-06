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
	var out output.DrawGacha

	// ガチャ対象キャラとレートタイプ取得
	// todo constをどこにかく？
	getRateInput.GachaType = 1
	getRateOutput, err := gi.gs.GetGachaRate(ctx, &getRateInput)
	if err != nil {
		return out, err
	}

	// ガチャ
	var drawGachaInput sInput.DrawGacha
	var charaIDs []int
	drawGachaInput.CharaRates = getRateOutput.CharaRates
	drawGachaInput.RateTypes = getRateOutput.RateTypes
	for i := 0; i < int(gacha.Times); i++ {
		drawGachaOutput, err := gi.gs.DrawGacha(ctx, &drawGachaInput)
		if err != nil {
			return out, err
		}
		charaIDs = append(charaIDs, drawGachaOutput.CharaID)
	}

	// 登録
	var addUserCharaInput sInput.AddUserChara
	addUserCharaInput.CharaIDs = charaIDs
	addUserCharaInput.UserID = gacha.UserID
	_, err = gi.cs.AddUserChara(ctx, &addUserCharaInput)
	if err != nil {
		return out, err
	}

	// キャラクター名取得
	var getCharasInput sInput.GetCharas
	getCharasInput.IDs = charaIDs
	getCharaOutput, err := gi.cs.GetCharas(ctx, &getCharasInput)
	if err != nil {
		return out, err
	}

	for _, charaID := range charaIDs {
		var chara output.Chara
		chara.ID = charaID
		chara.Name = getCharaOutput.CharaName[chara.ID]
		out.Charas = append(out.Charas, chara)
	}

	return out, nil
}
