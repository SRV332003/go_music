package player

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"dhvani/filemanager"
)

var streamer beep.Streamer
var format beep.Format
var seeker beep.StreamSeekCloser
var loop bool
var playing bool

var i int

func Play(file string) error {

	err := changeStream(file)

	i = 1

	speaker.Clear()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Iterate(iterator))

	playing = true

	return err
}

func PausePlay() {

	if playing {
		speaker.Clear()
		i = 2
		playing = false
	} else {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		speaker.Play(beep.Iterate(iterator))
		playing = true
	}
}

func Resume() {
	// log.Println("Resuming", i)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Iterate(iterator))
}

func Skip(t int) {

	err := seeker.Seek(seeker.Position() + (t * format.SampleRate.N(time.Second)))
	if err != nil {
		log.Println(err)
	}

}

func changeStream(file string) (err error) {

	f, err := os.Open(file)
	if err != nil {
		return
	}
	streamer, format, err = mp3.Decode(f)
	if err != nil {
		return
	}
	seeker = streamer.(beep.StreamSeekCloser)
	return err

}

func iterator() beep.Streamer {

	// log.Println("Iterator called", i)
	if i == 2 {
		i = 0
		// log.Println("Paused restored")
		return streamer
	} else if i == 1 || loop {
		i = 0
		seeker.Seek(0)
		// log.Println("Looping")
		return streamer
	}

	song := filemanager.GetRandom()
	changeStream(song.Path)

	go PausePlay()
	go Resume()

	return streamer
}

func init() {
	loop = true
	playing = false
}

func Loop(l bool) {
	loop = l
}

func Next() {
	seeker.Seek(seeker.Len())
	// log.Println(seeker.Position(), seeker.Len())
}
