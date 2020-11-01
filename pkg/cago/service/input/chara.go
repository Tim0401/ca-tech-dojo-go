package input

type GetCharas struct {
	IDs []int32
}

type AddUserChara struct {
	UserID   int32
	CharaIDs []int32
}
