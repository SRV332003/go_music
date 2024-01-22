package handler

import (
	"fmt"
	"go_music/downloader"
	"go_music/filemanager"
	"go_music/player"

	"github.com/eiannone/keyboard"
)

func handleSearch(s string) {
	res := filemanager.Search(s)
	ytres := downloader.FetchSearch(s)

	for i, song := range res {
		fmt.Printf("%3d. %s...\n", i, song.Name[:min(30, len(song.Name))])
	}
	count := len(res)

	for i := 0; i < min(10-len(res), len(ytres)); i++ {
		fmt.Printf("%3d. %30s\t%20s...\n", count+i, ytres[i][1][:min(30, len(ytres[i][1]))], ytres[i][2])
	}

	var char rune
	fmt.Print("Enter your choice (0 to exit): ")
	char, _, _ = keyboard.GetSingleKey()
	if char == '\x00' {
		fmt.Println("-1")
		return
	}

	n := int(char - '0')
	fmt.Println(n)

	for n < 0 || n > min(len(res)+len(ytres), 10) {
		fmt.Println("\rInvalid choice!!")
		fmt.Print("Enter your choice (0 to exit): ")
		fmt.Scanln(&n)
	}

	if n < len(res) {

		err := player.Play(res[n])
		if err != nil {
			panic(err)
		}

	} else {

		fmt.Println("Downloading...")
		name, dest, err := downloader.Getfile(ytres[n-count][0])
		if err != nil {
			panic(err)
		}
		err = player.Play(filemanager.AddSong(name, dest))
		if err != nil {
			panic(err)
		}

	}

}
