package middangeard

// Ranking represents the game's ranking system.
type Ranking struct {
	Score uint16 `json:"score"`
	Rank  string `json:"rank"`
}

var Rankings []Ranking

// Rankings implements possible adventurer ranks.
// var Rankings = map[uint16]string{
// 	0:                   "Beginner",
// 	System.MaxScore / 2: "Amateur Adventurer",
// 	System.MaxScore:     "Master Adventurer",
// }

// Rankings implements possible adventurer ranks.
// var Rankings = []Ranking{
// 	{0, "Beginner"},
// 	{System.MaxScore / 2, "Amateur Adventurer"},
// 	{System.MaxScore, "Master Adventurer"},
// }
