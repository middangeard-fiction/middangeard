package middangeard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse implements a simple, extensible string parser.
func (g *Game) Parse() {
	// reader should be wrapped inside an _input() function.
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		for reader.Scan() {
			switch reader.Text() {
			default:
				// Split into tokens prior to synonymization
				tokens := g._tokenize(reader.Text())
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
			}
			fmt.Print(">>> ")
		}
	}
}

func (g *Game) _tokenize(cmd string) []string {
	if len(cmd) != 0 {
		tokens := strings.Fields(cmd)
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
