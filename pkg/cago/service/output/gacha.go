package output

import "ca-tech-dojo-go/pkg/cago/service/io"

type DrawGacha struct {
	CharaID int32
}

type GetGachaRate struct {
	CharaRates []io.CharaRate
}
