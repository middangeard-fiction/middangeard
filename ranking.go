package middangeard

// Ranking represents the game's ranking system.
type Ranking struct {
	Score uint16 `json:"score"`
	Rank  string `json:"rank"`
}
