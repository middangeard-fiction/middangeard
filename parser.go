package middangeard

import (
	"bufio"
	"fmt"
	"strings"
)

type parser struct {
	console *bufio.ReadWriter
	reader  string
}

// var r = strings.NewReader("Hello, Reader!")
// var console *bufio.ReadWriter
// var reader string

// Parse implements a simple, extensible string parser.
func (g *Game) Parse() {
	// reader should be wrapped inside an _input() function.
	for {
		fmt.Print(">>> ")
		if len(g.parser.reader) == 0 {
			g.parser.reader, _ = g.parser.console.ReadString('\n')
		}
		switch g.parser.reader {
		default:
			// Split into tokens prior to synonymization
			tokens := g._tokenize(g.parser.reader)
			if len(tokens) > 0 {
				v := tokens[0]
				s := g._synonymize(v)
				if n, _ := g.Verbs[s]; g.Verbs[s] != nil {
					if len(tokens) > 1 {
						o := tokens[1]
						n(o)
					} else {
						n()
					}
				} else {
					// all Println' should later be wrapped into an _output() function.
					fmt.Println("I beg your pardon?")
				}
			} else {
				fmt.Println("I beg your pardon?")
			}
			g.parser.reader = ""
		}
	}
}

func (g *Game) _tokenize(cmd string) []string {
	if len(cmd) != 0 {
		tokens := strings.Fields(cmd)
		fmt.Println(tokens)
		return tokens
	}
	return []string{" "}
}

func (g *Game) _synonymize(cmd string) string {
	for verb, aliases := range g.Synonyms {
		for _, alias := range aliases {
			if alias == cmd {
				return verb
			} else if i := strings.Index(cmd, verb+" "); i == 0 {
				return cmd
			} else if i := strings.Index(cmd, alias+" "); i == 0 {
				return strings.Replace(cmd, alias+" ", verb+" ", 1)
			}
		}
	}
	return cmd
}
