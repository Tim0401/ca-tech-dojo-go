package input

type GetCharaList struct {
	Characters []UserChara `json:"characters"`
}

type UserChara struct {
	ID      int    `json:"userCharacterID"`
	CharaID int    `json:"characterID"`
	Name    string `json:"name"`
}
