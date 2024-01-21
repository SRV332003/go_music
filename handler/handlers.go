package handler

import (
	"dhvani/filemanager"
	"dhvani/player"
	"fmt"
	"os"
	"os/exec"

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
		player.Resume()
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid argument")
		return
	}
	player.Play(filemanager.GetSongByID(i - 1).Path)

}

func advSearchHandler(args []string) {
	song, err := filemanager.AdvSearch()
	if err != nil {
		fmt.Println("No song found")
		return
	}
	player.Play(song.Path)
	fmt.Println("Playing", song.Name)
}

func randomSongHandler(args []string) {
	song := filemanager.GetRandom()
	fmt.Println("Playing", song.Name)
	player.Play(song.Path)
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
