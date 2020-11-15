package output

import "ca-tech-dojo-go/pkg/cago/service/io"

type DrawGacha struct {
	CharaID int
}

type GetGachaRate struct {
	CharaProbability map[string]*io.CharaProbability
	GroupProbability *io.GroupProbability
}
