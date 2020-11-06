package output

import "ca-tech-dojo-go/pkg/cago/service/io"

type DrawGacha struct {
	CharaID int
}

type GetGachaRate struct {
	CharaRates map[int]*io.CharaRates
	RateTypes  *io.RateTypes
}
