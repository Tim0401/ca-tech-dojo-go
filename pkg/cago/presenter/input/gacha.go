package input

type DrawGacha struct {
	Results []Chara `json:"results"`
}

type Chara struct {
	ID   int    `json:"characterID"`
	Name string `json:"name"`
}
