package handler

import (
	"dhvani/downloader"
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
		filemanager.GetSongByID(i - 1).Play()
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
	ytres := downloader.FetchSearch(s)

	for i, song := range res {
		fmt.Printf("%3d. %s...\n", i+1, song.Name[:min(30, len(song.Name))])
	}
	count := len(res)

	for i := 0; i < min(10-len(res), len(ytres)); i++ {
		fmt.Printf("%3d. %30s\t%20s...\n", count+i+1, ytres[i][1][:min(30, len(ytres[i][1]))], ytres[i][2])
	}

	var n int
	fmt.Print("Enter your choice (0 to exit): ")
	fmt.Scanln(&n)

	for n <= 0 || n > min(len(res)+len(ytres), 10) {
		if n == 0 {
			return
		}
		fmt.Println("\rInvalid choice!!")
		fmt.Print("Enter your choice (0 to exit): ")
		fmt.Scanln(&n)
	}

	if n-1 < len(res) {

		fmt.Println("Playing", res[n-1].Name)
		res[n-1].Play()

	} else {

		fmt.Println("Downloading...")
		name, dest, err := downloader.Getfile(ytres[n-count-1][0])
		if err != nil {
			panic(err)
		}
		filemanager.AddSong(name, dest).Play()
		fmt.Println("Playing", name)

	}

}

func clearScr() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
