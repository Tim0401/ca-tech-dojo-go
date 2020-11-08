package input

import "ca-tech-dojo-go/pkg/cago/service/io"

type DrawGacha struct {
	CharaRates map[int]*io.CharaRates
	RateTypes  *io.RateTypes
}

type GetGachaRate struct {
	GachaType int
}
