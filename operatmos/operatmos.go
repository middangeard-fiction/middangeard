package operatmos

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Audio represents an audio stream.
type Audio struct {
	Filename   string
	controller *beep.Ctrl
}

// Initialize Operatmos during import, so the consumer doesn't have to.
func init() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

// StartSound plays the contents of a digital sound file.
func (a *Audio) StartSound(loop int) {
	f, err := os.Open(a.Filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(loop, streamer), Paused: false}
	speaker.Play(ctrl)

	a.controller = ctrl
}

// StartMusic is essentially a wrapper for "StartSound". Will call it with "loop" automatically set to INFINITE.
func (a *Audio) StartMusic() {
	a.StartSound(-1)
}

// Pause pauses audio playback.
func (a *Audio) Pause() {
	speaker.Lock()
	a.controller.Paused = true
	speaker.Unlock()
}

// Resume resumes audio playback.
func (a *Audio) Resume() {
	speaker.Lock()
	a.controller.Paused = false
	speaker.Unlock()
}
