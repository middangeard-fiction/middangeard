package middangeard

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/InVisionApp/conjungo"
)

/* Merge built-in verbs/synonyms with game author defined ones.
   In later iterations of Middangeard, there may be a way to provide an exclusion bit per verb,
   so game authors may conveniently enable/disable built-in verbs/synonyms as needed.
*/
func (g *Game) syncVerbs() {
	builtin := Game{
		Verbs: map[string]func(...string){
			"help":  g.help,
			"look":  g.look,
			"quit":  g.quit,
			"score": g.showScore,
			"go":    g.walk,
		},
		Synonyms: map[string][]string{
			"quit": {"exit"},
			"go":   {"walk", "travel"},
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

func (g *Game) look(...string) {
	r := g._lookAhead()
	fmt.Println(_wordWrap(r.Description, 60))
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

// Returns new room.
func (g *Game) _lookAhead() *Room {
	if r := g.Rooms[g.Player.Location]; r != nil {
		return r
	}
	return &Room{Description: `You're standing in a void of nothingness. Obviously, the author of this adventure has yet to implement this room.`}
}
