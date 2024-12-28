package player

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"

	"github.com/SRV332003/go_music/filemanager"
	"github.com/SRV332003/go_music/recom"
)

var streamer beep.Streamer
var format beep.Format
var seeker beep.StreamSeekCloser
var loop bool
var volume *effects.Volume
var resampler *beep.Resampler
var ctrl *beep.Ctrl

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
	speaker.Play(volume)

	ChangeVolume(0)

	fmt.Println("Playing", song.ID, "\b..", song.Name)
	// filemanager.UpdateUsage(song)

	return err
}

func PausePlay() {
	speaker.Lock()
	if ctrl.Paused {
		ctrl.Paused = false
	} else {
		ctrl.Paused = true
	}
	speaker.Unlock()
}

func Skip(t int) {
	// log.Println("Skipping", t)
	speaker.Lock()
	targetSeek := min(seeker.Len()-1, seeker.Position()+(t*format.SampleRate.N(time.Second)))
	targetSeek = max(0, targetSeek)
	// log.Println(targetSeek, seeker.Len(), seeker.Position()+t*format.SampleRate.N(time.Second))
	err := seeker.Seek(targetSeek)
	if err != nil {
		log.Println(err)
	}
	speaker.Unlock()

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
	ctrl = &beep.Ctrl{Streamer: beep.Iterate(iterator), Paused: false}
	resampler = beep.ResampleRatio(4, 1, ctrl)
	prevVolume := 0.0
	if volume != nil {
		prevVolume = volume.Volume
	}
	volume = &effects.Volume{
		Streamer: resampler,
		Base:     2,
		Volume:   float64(prevVolume),
		Silent:   false,
	}
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
	fmt.Println("\rPlaying", song)
	fmt.Print(color.CyanString("dhvani > "))

	return streamer
}

func init() {
	loop = true
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

// Change volume of the song
// n: volume change factor (-1.0 to 1.0)
func ChangeVolume(n float64) {
	// log.Println("Changing volume to", volume.Volume+n)
	if volume.Volume+n >= 1 {
		// log.Println("Volume too high")
		n = 0
	}
	if volume.Volume+n < -6 {
		// log.Println("Volume too low")
		n = 0
	}
	speaker.Lock()
	volume.Volume += float64(n)
	speaker.Unlock()

}

// change speed of the song
// n: speed change factor (-1.0 to 1.0)
func ChangeSpeed(n float64) {
	speaker.Lock()
	resampler.SetRatio(resampler.Ratio() + n)
	speaker.Unlock()
}
