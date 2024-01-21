package handler

import (
	"fmt"
	"go_music/downloader"
	"go_music/filemanager"
	"go_music/player"
	"strings"
)

func HandleWordCommands(s string) {

	if len(s) == 0 {
		fmt.Println("\b")
		return
	}

	if s[0] == ':' {
		handleSearch(s[1:])

	} else if s[0] == '~' {

		name, dest, err := downloader.Getfile(s[1:])
		if err != nil {
			panic(err)
		}
		player.Play(filemanager.AddSong(name, dest))

	} else {
		command := strings.Split(s, " ")[0]
		command = strings.ToLower(command)
		command = strings.TrimSpace(command)
		args := strings.Split(s, " ")[1:]

		switch command {

		case "ls":
			lsHandler(args)
		case "play":
			playHandler(args)
		case "cd":
			cdHandler(args)
		case "skip":
			skipHandler(args)

		}
	}

}
