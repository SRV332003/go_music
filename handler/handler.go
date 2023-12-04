package handler

import (
	"dhvani/filemanager"
	"fmt"
	"log"
	"strings"
)

func HandleInput(s string) {

	s = strings.Split(s, "\n")[0]
	fmt.Println("Handling input", s)
	log.Println("Handling input")

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
	case "play":
		// play(args)
	case "pause":
		// pause()
	case "resume":
		// resume()
	case "skip":
		// skip(args)
	case "exit":
		// exit()
	case "search":
		// search(args)
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

}
