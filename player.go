package middangeard

import (
	"bytes"
	"fmt"
	"text/template"
)

// Player represents the player character.
type Player struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	Inventory   Items  `json:"inventory"`
	Score       uint16 `json:"score"`
}

// PickupItem adds a specified item to the player inventory.
// => Shortcut for Player.Inventory.Add.
func (p *Player) PickupItem(item *Item) {
	p.Inventory.Add(item)
}

// AwardPoints increases the player's score.
// the message parameter can be used for an optional, congratulatory text.
func (p *Player) AwardPoints(points uint16, message ...string) {
	p.Score += points

	if len(message) != 0 {
		fmt.Println(message[0])
	}
}

func (p *Player) listInventory(...string) {
	var text string
	var tpl bytes.Buffer

	// Temporary hack, since Output is a method of Game.
	var g *Game

	var fns = template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}
	t := template.Must(template.New("").Funcs(fns).Parse(itemListTemplate))
	t.Execute(&tpl, p.Inventory)

	text = tpl.String()

	if len(text) != 0 {
		fmt.Println()
		// fmt.Printf("You are carrying %v.", text)
		g.Output("You are carrying %v.", text)
		fmt.Println()
	} else {
		g.Output("You aren't carrying anything.")
	}
}
