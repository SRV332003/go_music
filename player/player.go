package player

import (
	// "log"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var streamer beep.Streamer
var format beep.Format
var seeker beep.StreamSeekCloser

func Play(file string) (beep.Streamer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	streamer, format, err = mp3.Decode(f)
	if err != nil {
		return nil, err
	}

	speaker.Clear()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Iterate(iterator))
	seeker = streamer.(beep.StreamSeekCloser)

	return streamer, err
}

func Pause() {
	speaker.Clear()
}

func Resume() {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

func Skip(t int) {

	// log.Println("Seeking from", seeker.Position()/format.SampleRate.N(time.Second), "to", seeker.Position()/format.SampleRate.N(time.Second)+t)

	err := seeker.Seek(seeker.Position() + (t * format.SampleRate.N(time.Second)))
	if err != nil {
		log.Println(err)
	}

}

func iterator() beep.Streamer {
	seeker.Seek(0)
	return streamer
}
