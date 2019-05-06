package middangeard

// Room represents a room in the game.
type Room struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image,omitempty"`
	Sound       string      `json:"sound,omitempty"`
	OnEnter     func(*Room) `json:"-"`
	OnLeave     func(*Room) `json:"-"`
	// Potentially, move things like "Lit" to a substructure "Attributes"
	Lit        bool       `json:"lit"`
	Visited    bool       `json:"visited"`
	Directions Directions `json:"directions"`
}

type Directions struct {
	North     string `json:"north,omitempty"`
	Northeast string `json:"northeast,omitempty"`
	Northwest string `json:"northwest,omitempty"`
	South     string `json:"south,omitempty"`
	Southeast string `json:"southeast,omitempty"`
	Southwest string `json:"southwest,omitempty"`
	East      string `json:"east,omitempty"`
	West      string `json:"west,omitempty"`
	Up        string `json:"up,omitempty"`
	Down      string `json:"down,omitempty"`
	In        string `json:"in,omitempty"`
	Out       string `json:"out,omitempty"`
}

// AddItem adds an item to a room. Potentially deprecated, as Middangeard should be declarative in nature.
func (p *Room) AddItem() {

}
