package middangeard

type mode int

type display struct {
	Console mode
	GUI     mode
}

// Display provides the possible display options
var Display = &display{
	Console: 0,
	GUI:     1,
}
