package input

type GetUserRanking struct {
	Results []UserRank `json:"results"`
}

type UserRank struct {
	Rank     int    `json:"rank"`
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
	Score    int    `json:"score"`
}
