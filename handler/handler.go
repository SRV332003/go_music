package handler

import (
	"dhvani/downloader"
	"dhvani/filemanager"
	"dhvani/player"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func HandleInput(s string) {

	s = strings.Split(s, "\n")[0]

	if s == "" {
		fmt.Print("\b")
		return
	}

	if s[0] == ':' {

		handleSearch(s[1:])

	} else if s[0] == '~' {

		name, dest, err := downloader.Getfile(s[1:])
		if err != nil {
			panic(err)
		}

		player.Play(filemanager.AddSong(name, dest).Path)
		fmt.Println("Playing", name)

	} else {

		s = strings.ToLower(s)
		command := strings.Split(s, " ")[0]
		args := strings.Split(s, " ")[1:]
		HandleCommand(command, args)

	}

}

func HandleCommand(command string, args []string) {

	switch command {

	case "clear":
		ClearScr()
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
		player.Play(filemanager.GetSongByID(i - 1).Path)

	case "p":
		player.PausePlay()
	case "resume":
		player.PausePlay()
	case "r":
		song := filemanager.GetRandom()
		fmt.Println("Playing", song.Name)
		player.Play(song.Path)
	case "skip":
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
		fmt.Println("Bye Bye! Miss You :')")
		os.Exit(0)
	case "help":
		// help()

	case "loop":
		if len(args) == 0 || args[0] == "on" {
			player.Loop(true)
			fmt.Println("Okay, I'll repeat it :)")
			return
		}
		if args[0] == "off" {
			player.Loop(false)
			fmt.Println("Okay, No repitions now :)")
			return
		}
	case "next":
		player.Next()
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
		player.Play(res[n-1].Path)

	} else {

		fmt.Println("Downloading...")
		name, dest, err := downloader.Getfile(ytres[n-count-1][0])
		if err != nil {
			panic(err)
		}
		player.Play(filemanager.AddSong(name, dest).Path)
		fmt.Println("Playing", name)

	}

}

func ClearScr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
