package middangeard

// Character represents any other character in the game.
type Character struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Loot        []string `json:"loot"`
}
