package io

type CharaRate struct {
	CharaID                 int
	GachaProbabilityGroupID string
	Rate                    int
}
type GachaProbabilityGroup struct {
	ID   string
	Rate int
}

type CharaProbability struct {
	CharaRates []*CharaRate
	SumRate    int
}
type GroupProbability struct {
	GachaProbabilityGroups []*GachaProbabilityGroup
	SumRate                int
}
