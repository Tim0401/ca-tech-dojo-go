package output

type GetCharaList struct {
	Charas []UserChara
}

type UserChara struct {
	ID      int
	CharaID int
	Name    string
}
