package handler

import (
	"dhvani/downloader"
	"dhvani/filemanager"
	"dhvani/player"
	"fmt"
	"strings"
)

func HandleWordCommands(s string) {

	if len(s) == 0 {
		fmt.Println("\b")
		return
	}

	switch s[0] {

	case ':':
		handleSearch(s[1:])
	case '~':
		name, dest, err := downloader.Getfile(s[1:])
		if err != nil {
			panic(err)
		}
		player.Play(filemanager.AddSong(name, dest).Path)
		fmt.Println("Playing", name)

	default:
		command := strings.Split(s, " ")[0]
		command = strings.ToLower(command)
		args := strings.Split(s, " ")[1:]

		switch command {

		case "ls":
			lsHandler(args)
		case "play":
			playHandler(args)
		case "resume":
			pausePlayHandler(args)
		case "cd":
			cdHandler(args)
		case "skip":
			skipHandler(args)
		}
	}

}
