package middangeard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/VladimirMarkelov/clui"
	term "github.com/nsf/termbox-go"
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

// ParseTerm implements an experimental TUI.
func (g *Game) ParseTerm() {
	var dialog *clui.ConfirmationDialog

	clui.InitLibrary()
	defer clui.DeinitLibrary()

	view := clui.AddWindow(0, 0, 80, 20, g.Title)
	view.SetTitleButtons(clui.ButtonMaximize | clui.ButtonClose)

	frmLeft := clui.CreateFrame(view, 8, 4, clui.BorderNone, 1)
	frmLeft.SetPack(clui.Vertical)
	frmLeft.SetGaps(clui.KeepValue, 1)
	frmLeft.SetPaddings(1, 1)
	logBox := clui.CreateTextView(frmLeft, 28, 50, clui.Fixed)
	logBox.SetWordWrap(true)
	logBox.SetMaxItems(0)

	frmTheme := clui.CreateFrame(frmLeft, 8, 1, clui.BorderNone, clui.Fixed)
	frmTheme.SetGaps(1, clui.KeepValue)
	clui.CreateFrame(frmLeft, 1, 1, clui.BorderNone, 1)

	input := clui.CreateEditField(frmLeft, 80, "", clui.Fixed)
	clui.ActivateControl(view, input)

	frmRight := clui.CreateFrame(view, 8, 4, clui.BorderNone, 1)
	frmRight.SetPack(clui.Vertical)
	frmRight.SetGaps(clui.KeepValue, 1)
	frmRight.SetPaddings(1, 1)

	frmPb := clui.CreateFrame(frmRight, 8, 1, clui.BorderNone, clui.Fixed)
	t := time.Now().Format("3:04PM")
	clui.CreateLabel(frmPb, 10, 1, t, clui.Fixed)

	frmTheme2 := clui.CreateFrame(frmRight, 8, 1, clui.BorderNone, clui.Fixed)
	frmTheme2.SetGaps(1, clui.KeepValue)
	clui.CreateFrame(frmRight, 1, 1, clui.BorderNone, 1)

	frmPb2 := clui.CreateFrame(frmRight, 8, 1, clui.BorderNone, clui.Fixed)
	clui.CreateLabel(frmPb2, 1, 1, "[", clui.Fixed)
	pb2 := clui.CreateProgressBar(frmPb2, 20, 1, 1)
	pb2.SetLimits(0, int(g.MaxScore))
	pb2.SetTitle("Score: {{value}} of {{max}}")
	clui.CreateLabel(frmPb2, 1, 1, "]", clui.Fixed)

	// Actual parser logic starts here.
	logBox.AddText([]string{g.Intro})
	input.OnKeyPress(func(key term.Key, ch rune) bool {
		if key == term.KeyCtrlM {
			switch input.Title() {
			case "quit":
				if len(g.Goodbye) != 0 {
					dialog = clui.CreateAlertDialog("Goodbye", g.Goodbye, "Exit")
				} else {
					dialog = clui.CreateAlertDialog("Goodbye", "Exiting game", "Exit")
				}
				dialog.OnClose(func() {
					clui.Stop()
				})
			default:
				for n, c := range g.Verbs {
					switch input.Title() {
					case n:
						c()
					default:
						logBox.AddText([]string{"\n" + "<c:yellow>I beg your pardon?"})
					}
				}
			}
			input.Clear()
			return true
		}
		return false
	})

	clui.MainLoop()
}
