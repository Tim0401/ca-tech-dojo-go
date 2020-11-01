package output

type DrawGacha struct {
	CharaRates []CharaRate
	Times      int32
}
type GetGachaRate struct {
	CharaRates []CharaRate
}

type CharaRate struct {
	CharaID int32
	Rate    float64
}
