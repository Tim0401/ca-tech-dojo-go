package input

import "ca-tech-dojo-go/pkg/cago/service/io"

type DrawGacha struct {
	CharaProbability map[string]*io.CharaProbability
	GroupProbability *io.GroupProbability
}

type GetGachaRate struct {
	GachaType int
}
