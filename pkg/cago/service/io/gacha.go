package io

type CharaRate struct {
	CharaID    int
	RateTypeID int
	Rate       int
}
type RateType struct {
	ID   int
	Rate int
}

type CharaRates struct {
	CharaRateArray []*CharaRate
	SumRate        int
}
type RateTypes struct {
	RateTypeArray []*RateType
	SumRate       int
}
