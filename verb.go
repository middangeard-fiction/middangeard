package middangeard

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/InVisionApp/conjungo"
)

// ItemVerbs represents Item Verb functions.
type ItemVerbs map[string]func(item *Item, room *Room)

/* Merge built-in verbs/synonyms with game author defined ones.
   In later iterations of Middangeard, there may be a way to provide an exclusion bit per verb,
   so game authors may conveniently enable/disable built-in verbs/synonyms as needed.
*/
func (g *Game) syncVerbs() {
	builtin := Game{
		Verbs: map[string]func(...string){
			"help":      g.help,
			"drop":      g.drop,
			"take":      g.take,
			"inventory": g.Player.listInventory,
			"inspect":   g.inspect,
			"look":      g.look,
			"quit":      g.quit,
			"score":     g.showScore,
			"go":        g.walk,
		},
		Synonyms: map[string][]string{
			"inventory": {"i"},
			"quit":      {"exit"},
			"go":        {"walk", "travel"},
		},
	}

	err := conjungo.Merge(&g.Synonyms, builtin.Synonyms, nil)
	if err != nil {
		fmt.Println(err)
	}

	err2 := conjungo.Merge(&g.Verbs, builtin.Verbs, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	// fmt.Println(g.Verbs)

	// b, err := json.Marshal(g.Synonyms)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))
}

func (g *Game) help(...string) {
	if len(g.Help) != 0 {
		g.Output(g.Help)
	} else {
		g.Output("In order to progress, type commands.")
	}
}

func (g *Game) showScore(...string) {
	g.Output("Current Score: %v out of %v\n", g.Player.Score, g.MaxScore)
}

func (g *Game) quit(object ...string) {
	// Check whether player actually wants to quit the game or quit drinking ;)
	if len(object) == 0 {
		if len(g.Goodbye) != 0 {
			g.Output(g.Goodbye)
		} else {
			// fmt.Println("Goodbye!")
			g.Output("Goodbye!")
		}
		os.Exit(0)
	}
}

func (g *Game) drop(args ...string) {
	if len(args) != 0 {
		obj := g.Player.Inventory._findObject(args)
		if obj != nil && len(args) != 0 {
			if r := g.Rooms[g.Player.Location]; r != nil {
				obj.Fixture = true
				r.Items.Add(obj)
				if cb := obj.Verbs["drop"]; cb != nil {
					cb(obj, r)
				}
			}
			g.Player.Inventory.Remove(obj)
			// Needs some sort of override, in case we don't actually allow to drop the item
			// See Cloak of Darkness cloak, for starters.
			// g.Output("You drop %v %v.", obj.Article, obj.Name)
		} else {
			g.Output("You aren't carrying any %v.", args[0])
			fmt.Println()
		}
	} else {
		// fmt.Printf("Drop what?")
		g.Output("Drop what?")
		fmt.Println()
	}
}

func (g *Game) take(args ...string) {
	if len(args) != 0 {
		obj := g.Rooms[g.Player.Location]._findObject(args)
		if obj != nil && len(args) != 0 {
			if obj.Carryable {
				if r := g.Rooms[g.Player.Location]; r != nil {
					r.Items.Remove(obj)
					if cb := obj.Verbs["take"]; cb != nil {
						cb(obj, r)
					}
				}
				g.Player.PickupItem(obj)
			} else {
				g.Output("You can't take %v.", obj.Name)
			}
		} else {
			g.Output("You can't see any %v.", args[0])
			fmt.Println()
		}
	} else {
		g.Output("Take what?")
		fmt.Println()
	}
}

func (g *Game) inspect(args ...string) {
	if len(args) != 0 {
		obj := g._findObject(args)
		if obj != nil && len(args) != 0 {
			if r := g.Rooms[g.Player.Location]; r != nil {
				if cb := obj.Verbs["inspect"]; cb != nil {
					cb(obj, r)
				}
				if len(obj.Description) != 0 {
					// fmt.Println(_wordWrap(obj.Description, 60))
					g.Output(obj.Description)
				}
			}
		} else {
			g.Output("You can't see any %v.", args[0])
			fmt.Println()
		}
	} else {
		g.Output("Look at what?")
		fmt.Println()
	}
}

func (g *Game) look(...string) {
	r := g._lookAhead()
	// fmt.Println(_wordWrap(r.Description, 60))
	g.Output(r.Description)

	var text string
	var tpl bytes.Buffer

	var fns = template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}

	t := template.Must(template.New("").Funcs(fns).Parse(itemRoomListTemplate))
	t.Execute(&tpl, r.Items)

	text = tpl.String()
	if len(text) != 0 {
		fmt.Println()
		g.Output("You see %v.", text)
		fmt.Println()
	}
}

func (g *Game) walk(direction ...string) {
	if len(direction) != 0 {
		switch direction[0] {
		case "north":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.North] != nil {
				g.Player.Location = r.Directions.North
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.North] == nil && r.Directions.North != "" {
				g.Output(r.Directions.North)
			}
		case "northeast":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Northeast] != nil {
				g.Player.Location = r.Directions.Northeast
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Northeast] == nil && r.Directions.Northeast != "" {
				g.Output(r.Directions.Northeast)
			}
		case "northwest":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Northwest] != nil {
				g.Player.Location = r.Directions.Northwest
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Northwest] == nil && r.Directions.Northwest != "" {
				g.Output(r.Directions.Northwest)
			}
		case "south":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.South] != nil {
				g.Player.Location = r.Directions.South
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.South] == nil && r.Directions.South != "" {
				g.Output(r.Directions.South)
			}
		case "southeast":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Southeast] != nil {
				g.Player.Location = r.Directions.Southeast
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Southeast] == nil && r.Directions.Southeast != "" {
				g.Output(r.Directions.Southeast)
			}
		case "southwest":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Southwest] != nil {
				g.Player.Location = r.Directions.Southwest
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Southwest] == nil && r.Directions.Southwest != "" {
				g.Output(r.Directions.Southwest)
			}
		case "east":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.East] != nil {
				g.Player.Location = r.Directions.East
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.East] == nil && r.Directions.East != "" {
				g.Output(r.Directions.East)
			}
		case "west":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.West] != nil {
				g.Player.Location = r.Directions.West
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.West] == nil && r.Directions.West != "" {
				g.Output(r.Directions.West)
			}
		case "up":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Up] != nil {
				g.Player.Location = r.Directions.Up
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Up] == nil && r.Directions.Up != "" {
				g.Output(r.Directions.Up)
			}
		case "down":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Down] != nil {
				g.Player.Location = r.Directions.Down
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Down] == nil && r.Directions.Down != "" {
				g.Output(r.Directions.Down)
			}
		case "in":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.In] != nil {
				g.Player.Location = r.Directions.In
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.In] == nil && r.Directions.In != "" {
				g.Output(r.Directions.In)
			}
		case "out":
			if r := g.Rooms[g.Player.Location]; g.Rooms[r.Directions.Out] != nil {
				g.Player.Location = r.Directions.Out
				fmt.Println()
				r.Visited = true
				if r.OnLeave != nil {
					r.OnLeave(r)
				}
				newRoom := g._lookAhead()
				if newRoom.OnEnter != nil {
					newRoom.OnEnter(newRoom)
				}
				g.look()
			} else if g.Rooms[r.Directions.Out] == nil && r.Directions.Out != "" {
				g.Output(r.Directions.Out)
			}
		}
	} else {
		g.Output("Walk where?")
	}
}

// Invoke sends an action (that would normally be performed by the player) to the parser.
func (g *Game) Invoke(action string) {
	g.parser.reader = action
	g.Parse()
}
