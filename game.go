package middangeard

import (
	"fmt"

	"github.com/Happy-Ferret/middangeard/parser"
)

// Game represents the game's ruleset.
type Game struct {
	Title    string `json:"title"`
	Intro    string `json:"intro"`
	Outro    string `json:"outro"`
	GameOver string `json:"gameover"`
	MaxScore uint16 `json:"maxScore"`
}

// NewGame initializes a new game, prints the intro and launches
// the text parser.
func (g *Game) NewGame() {
	fmt.Println(g.Intro)

	parser.Parse()
}
