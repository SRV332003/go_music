package player

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"go_music/filemanager"
	"go_music/recom"
)

var streamer beep.Streamer
var format beep.Format
var seeker beep.StreamSeekCloser
var loop bool
var playing bool

var volume = 0

var i int

func Play(song filemanager.Song) error {

	if song.ID == 0 {
		fmt.Println("Song not found")
		return fmt.Errorf("song not found")
	}

	err := changeStream(song.Path)

	i = 1

	speaker.Clear()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Iterate(iterator))

	ChangeVolume(0)

	fmt.Println("Playing", song.ID, "\b..", song.Name)
	// filemanager.UpdateUsage(song)

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
	targetSeek := min(seeker.Len()-1, seeker.Position()+(t*format.SampleRate.N(time.Second)))
	// log.Println(targetSeek, seeker.Len(), seeker.Position()+t*format.SampleRate.N(time.Second))
	err := seeker.Seek(targetSeek)
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

	song := recom.PlayRandom()
	changeStream(song.Path)
	fmt.Print("\rPlaying", song)
	fmt.Print("\ndhvani > ")

	go PausePlay()
	go Resume()

	return streamer
}

func init() {
	loop = true
	playing = false
}

func SetLoop(l bool) {
	loop = l
}

func GetLoop() bool {
	return loop
}

func Next() {
	seeker.Seek(seeker.Len())
	// log.Println(seeker.Position(), seeker.Len())
}

func ChangeVolume(n int) {
	speaker.Lock()

	volume = volume + n
	if volume < -10 {
		volume = -10
	}

	if volume > 0 {
		volume = 0
	}
	i = 2

	ctrl := &beep.Ctrl{Streamer: beep.Iterate(iterator), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   float64(volume),
		Silent:   false,
	}
	speaker.Unlock()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(volume)

}
