package middangeard

import (
	"bytes"
	"encoding/json"
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
	fmt.Println(g.Verbs)

	b, err := json.Marshal(g.Synonyms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func (g *Game) help(...string) {
	if len(g.Help) != 0 {
		fmt.Println(g.Help)
	} else {
		fmt.Println("In order to progress, type commands")
	}
}

func (g *Game) showScore(...string) {
	fmt.Printf("Current Score: %v out of %v\n", g.Player.Score, g.MaxScore)
}

func (g *Game) quit(object ...string) {
	// Check whether player actually wants to quit the game or quit drinking ;)
	if len(object) == 0 {
		if len(g.Goodbye) != 0 {
			fmt.Println(g.Goodbye)
		} else {
			fmt.Println("Goodbye!")
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
		} else {
			fmt.Printf("You aren't carrying any %v.", args[0])
			fmt.Println()
		}
	} else {
		fmt.Printf("Drop what?")
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
				fmt.Printf("You can't take %v", obj.Name)
			}
		} else {
			fmt.Printf("You can't see any %v.", args[0])
			fmt.Println()
		}
	} else {
		fmt.Printf("Take what?")
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
					fmt.Println(_wordWrap(obj.Description, 60))
				}
			}
		} else {
			fmt.Printf("You can't see any %v.", args[0])
			fmt.Println()
		}
	} else {
		fmt.Printf("Look at what?")
		fmt.Println()
	}
}

func (g *Game) look(...string) {
	r := g._lookAhead()
	fmt.Println(_wordWrap(r.Description, 60))

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
		fmt.Printf("You see %v.", text)
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
				fmt.Println(_wordWrap(r.Directions.North, 60))
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
				fmt.Println(_wordWrap(r.Directions.Northeast, 60))
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
				fmt.Println(_wordWrap(r.Directions.Northwest, 60))
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
				fmt.Println(_wordWrap(r.Directions.South, 60))
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
				fmt.Println(_wordWrap(r.Directions.Southeast, 60))
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
				fmt.Println(_wordWrap(r.Directions.Southwest, 60))
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
				fmt.Println(_wordWrap(r.Directions.East, 60))
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
				fmt.Println(_wordWrap(r.Directions.West, 60))
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
				fmt.Println(_wordWrap(r.Directions.Up, 60))
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
				fmt.Println(_wordWrap(r.Directions.Down, 60))
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
				fmt.Println(_wordWrap(r.Directions.In, 60))
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
				fmt.Println(_wordWrap(r.Directions.Out, 60))
			}
		}
	} else {
		fmt.Println("Walk where?")
	}
}

// Invoke sends an action (that would normally be performed by the player) to the parser.
func (g *Game) Invoke(action string) {
}
