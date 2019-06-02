package middangeard

import (
	"fmt"
)

type Item struct {
	Article     string    `json:"article"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Verbs       ItemVerbs `json:"-"`
	Synonyms    []string  `json:"synonyms"`
	Moveable    bool      `json:"moveable"`
	Carryable   bool      `json:"carryable"`
	Attributes  []string  `json:"attributes"`
	Fixture     bool      `json:"fixture,omitempty"`
}

// Items is a slice of Item pointers.
type Items []*Item

const itemListTemplate = `{{$first := true}}{{$last := len .}}{{range $key, $value := $}}{{if $first}}{{$first = false}}{{else if ne (plus1 $key) $last}}, {{else if eq (plus1 $key) $last}} and {{end}}{{$value.Article}} {{$value.Name}}{{end}}`
const itemRoomListTemplate = `{{$first := true}}{{$last := len .}}{{range $key, $value := $}}{{if $value.Fixture}}{{if $first}}{{$first = false}}{{else if ne (plus1 $key) $last}}, {{else if eq (plus1 $key) $last}} and {{end}}{{$value.Article}} {{$value.Name}}{{end}}{{end}}`

// // Attributes = []mid.Attribute{
// 		{0, "Neatness"},
// 	}

// Potentially, an object map instead.

// Add a specified item to the item container.
func (i *Items) Add(item *Item) {
	*i = append(*i, item)
}

// Remove a specified item from the item container.
func (i *Items) Remove(item *Item) {
	loc := -1
	for i, val := range *i {
		if val == item {
			loc = i
			fmt.Println(i)
			break
		}
	}
	*i = append((*i)[:loc], (*i)[loc+1:]...)
}

func (i *Item) _lookUp(args []string) bool {
	if args[0] == i.Name {
		return true
	}
	for _, alias := range i.Synonyms {
		if args[0] == alias {
			return true
		}
	}
	return false
}

func (i *Items) _findObject(args []string) *Item {
	for _, obj := range *i {
		if obj._lookUp(args) {
			return obj
		}
	}
	return nil
}
