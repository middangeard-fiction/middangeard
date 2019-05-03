package middangeard

type mode int

type display struct {
	Console mode
	TUI     mode
	GUI     mode
}

// Display provides the possible display options
var Display = &display{
	Console: 0,
	TUI:     1,
	GUI:     2,
}
