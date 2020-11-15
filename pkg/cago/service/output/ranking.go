package output

type GetUserRanking struct {
	Ranks []UserRank
}

type UserRank struct {
	Rank   int
	UserID int
	Score  int
}
