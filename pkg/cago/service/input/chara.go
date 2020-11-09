package input

type GetCharas struct {
	IDs []int
}

type AddUserChara struct {
	UserID   int
	CharaIDs []int
}

type GetUserCharas struct {
	UserID int
}
