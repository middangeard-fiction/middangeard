package middangeard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// Message is an incoming IPC message.
type Message struct {
	Name    string          `json:"name"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

var socket *websocket.Conn

// MessageOut is an outgoing IPC message.
type MessageOut struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload,omitempty"`
}

// Collection of connected clients.
var clients = make(map[*websocket.Conn]bool)

func (g *Game) initSocket() {
	// fmt.Println("Start Listening")
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 2)
			// fmt.Fprintln(w, "")
		})
		// Final name for io system? Fafnir?
		http.Handle("/fafnir", websocket.Handler(g.messageHandler))

		http.ListenAndServe(":8088", nil)
	}()
}

func (g *Game) messageHandler(ws *websocket.Conn) {
	clients[socket] = true
	socket = ws

	for {
		var m Message
		if err := websocket.JSON.Receive(socket, &m); err != nil {
			fmt.Printf("error: %v", err)
			delete(clients, socket)
		}

		switch m.Name {
		case "launch":
			fmt.Println("GUI launched")
			// g.Parse()
		case "input":
			// message, _ := json.Marshal(m)
			// s := string(message)

			var p string
			json.Unmarshal(m.Payload, &p)

			g.parser.reader = p
			g.Parse()

			// message, _ := json.Marshal(&MessageOut{Name: m.Name + ".callback", Payload: id})
			// out := string(message)
			// websocket.Message.Send(ws, out)
		}
	}
}

// Make part of IO struct or a "Text" struct. OR make regular, global function (mid.Output).
func (g *Game) Output(text string, a ...interface{}) {
	fText := fmt.Sprintf(text, a...)

	switch displayMode {
	case Display.Console:
		fmt.Println(_wordWrap(fText, 60))

		if socket != nil {
			message, _ := json.Marshal(&MessageOut{Name: "output", Payload: fText})
			out := string(message)
			websocket.Message.Send(socket, out)
		}
	case Display.GUI:
		if socket != nil {
			// Shorten "input/output" to "in/out"?
			message, _ := json.Marshal(&MessageOut{Name: "output", Payload: fText})
			out := string(message)

			websocket.Message.Send(socket, out)
		}
	}
}
