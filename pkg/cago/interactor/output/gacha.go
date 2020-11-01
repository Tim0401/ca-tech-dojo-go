package output

type DrawGacha struct {
	Charas []chara
}

type chara struct {
	ID   int32
	Name string
}
