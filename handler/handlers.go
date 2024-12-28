package handler

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/SRV332003/go_music/filemanager"
	"github.com/SRV332003/go_music/player"

	"github.com/eiannone/keyboard"

	"strconv"
)

func loopHandler(args []string) {
	if player.GetLoop() {
		player.SetLoop(false)
		fmt.Println("Okay, I won't repeat it :)")
		return
	} else {
		player.SetLoop(true)
		fmt.Println("Okay, I'll repeat it :)")
		return
	}
}

func skipHandler(args []string) {
	if len(args) == 0 {
		player.Skip(10)
		return
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid argument")
		return
	}
	player.Skip(i)
}

func playHandler(args []string) {
	if len(args) == 0 {
		player.PausePlay()
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid argument")
		return
	}
	player.Play(filemanager.GetSongByID(i - 1))

}

func advSearchHandler(args []string) {
	song, err := filemanager.AdvSearch()
	if err != nil {
		fmt.Println("No song found")
		return
	}
	player.Play(song)
}

func randomSongHandler(args []string) {
	song := filemanager.GetRandom()
	player.Play(song)
}

func exitHandler(args []string) {
	fmt.Println("Bye Bye! Miss You :')")
	os.Exit(0)
}

func cdHandler(args []string) {
	if len(args) == 0 {
		fmt.Println("Invalid argument")
		return
	}

	filemanager.SetMusicDir(args[0])
}

func nextHandler(args []string) {
	player.Next()
}

func lsHandler(args []string) {
	filemanager.ListFiles()
}

func pausePlayHandler(args []string) {
	player.PausePlay()
}

func ClearScr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func volumeUpHandler(args []string) {
	player.ChangeVolume(0.5)
}

func volumeDownHandler(args []string) {
	player.ChangeVolume(-0.5)
}

func speedUpHandler(args []string) {
	player.ChangeSpeed(0.1)
}

func speedDownHandler(args []string) {
	player.ChangeSpeed(-0.1)
}

func help() {
	ClearScr()
	fmt.Printf(`For basic usage you can refer the following commands:

r -> random song in home music directory ~/Music
l -> toggle the loop
n -> next song
s -> search and play song in music directory
p -> toggle play, pause
q -> exit the player
c -> clear the terminal
h -> show help commands
: -> allows to write word commands.

Words commands:

cd <full-directory-path> --> changes the music directory and reloads the files
: <song-name>            --> show search results of youtube and allow playing/downloading from the results
~ <youtube-song-link>    --> download the audio of the link in music directory and play
skip <second>            --> skip the current song by given seconds (default = 10sec)
ls                       --> shows all songs in music directory with ids
play <id>                --> play the song with given id

Press any key to continue...`)
	keyboard.GetSingleKey()
	ClearScr()
}
