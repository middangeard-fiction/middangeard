package middangeard

// Player represents the player character.
type Player struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Inventory   []string `json:"inventory"`
	Score       uint16   `json:"score"`
}
