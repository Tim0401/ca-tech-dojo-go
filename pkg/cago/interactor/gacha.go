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
}

// NewGachaInteractor NewGachaInteractor
func NewGachaInteractor(gs service.GachaService) GachaInteractor {
	return &gachaInteractor{gs}
}

func (gi *gachaInteractor) DrawGacha(ctx context.Context, gacha *input.DrawGacha) (output.DrawGacha, error) {

	var getRateInput sInput.GetGachaRate
	var output output.DrawGacha

	_, err := gi.gs.GetGachaRate(ctx, &getRateInput)
	if err != nil {
		return output, err
	}

	return output, nil
}
