package middangeard

// Player represents the player character.
type Player struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location"`
	Inventory   []string `json:"inventory"`
	Score       uint16   `json:"score"`
}
