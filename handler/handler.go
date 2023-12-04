package handler

import (
	"dhvani/filemanager"
	"dhvani/player"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func HandleInput(s string) {

	s = strings.Split(s, "\n")[0]
	s = strings.ToLower(s)

	if s == "" {
		fmt.Print("\b")
		return
	}

	if s[0] == ':' {

		handleSearch(s[1:])

	} else {

		command := strings.Split(s, " ")[0]
		args := strings.Split(s, " ")[1:]
		HandleCommand(command, args)

	}

}

func HandleCommand(command string, args []string) {

	switch command {

	case "clear":
		clearScr()
	case "ls":
		filemanager.ListFiles()
	case "play":
		if len(args) == 0 {
			player.Resume()
			return
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid argument")
			return
		}
		filemanager.GetSongByID(i).Play()
	case "pause":
		player.Pause()
	case "resume":
		player.Resume()
	case "skip":
		log.Println("Skipping", args)
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
	case "exit":
		os.Exit(0)
	case "help":
		// help()
	default:
		// help()
	}
}

func handleSearch(s string) {
	res := filemanager.Search(s)

	for i, song := range res {
		fmt.Printf("%3d. %s...\n", i+1, song.Name[:min(30, len(song.Name))])
	}

	var n int
	fmt.Print("Enter your choice (0 to exit): ")
	fmt.Scanln(&n)

	for n <= 0 || n > len(res) {
		if n == 0 {
			return
		}
		fmt.Println("\rInvalid choice!!")
		fmt.Print("Enter your choice (0 to exit): ")
		fmt.Scanln(&n)
	}

	fmt.Println("Playing", res[n-1].Name)
	res[n-1].Play()

}

func clearScr() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
