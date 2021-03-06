package middangeard

import (
	"fmt"
	"strings"

	"github.com/middangeard-fiction/middangeard/operatmos"
)

// Game represents the game's ruleset.
type Game struct {
	Title    string                     `json:"title"`
	Author   string                     `json:"author"`
	Intro    string                     `json:"intro"`
	Outro    string                     `json:"outro"`
	Goodbye  string                     `json:"goodbye"`
	GameOver string                     `json:"gameover"`
	Help     string                     `json:"help"`
	MaxScore uint16                     `json:"maxScore"`
	Verbs    map[string]func(...string) `json:"-"`
	Synonyms map[string][]string        `json:"synonyms"`
	Player   Player
	Items    map[*Item]*Item
	Rooms    map[string]*Room `json:"-"`
	Sounds   map[*operatmos.Audio]*operatmos.Audio
}

// NewGame initializes a new game, prints the intro and launches
// the text parser.
func (g *Game) NewGame(d mode) {
	if Rankings == nil {
		Rankings = []Ranking{
			{0, "Beginner"},
			{g.MaxScore / 2, "Amateur Adventurer"},
			{g.MaxScore, "Master Adventurer"},
		}
	}

	g.syncVerbs()

	fmt.Printf("%v by %v \n \n", g.Title, g.Author)
	fmt.Println(_wordWrap(g.Intro, 60))

	for id, room := range g.Rooms {
		switch id {
		case g.Player.Location:
			fmt.Println()
			fmt.Println(_wordWrap(room.Description, 60))
		}
	}

	switch d {
	case Display.Console:
		g.Parse()
	case Display.GUI:
		fmt.Println("Display Mode TBI")
	}
}

func _wordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}

// Look up objects in current room AND the player inventory.
// TODO: Support for items serving as containers for other items.
func (g *Game) _findObject(args []string) *Item {
	if r := g.Rooms[g.Player.Location]; r != nil {
		if obj := r.Items._findObject(args); obj != nil {
			return obj
		}
	}
	if obj := g.Player.Inventory._findObject(args); obj != nil {
		return obj
	}
	return nil
}

// Returns new room.
func (g *Game) _lookAhead() *Room {
	if r := g.Rooms[g.Player.Location]; r != nil {
		return r
	}
	return &Room{Description: `You're standing in a void of nothingness. Obviously, the author of this adventure has yet to implement this room.`}
}
