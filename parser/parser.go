package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse implements a simple, extensible string parser.
func Parse() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "quit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("I beg your pardon?")
		}
	}
}
