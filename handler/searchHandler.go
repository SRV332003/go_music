package handler

import (
	"dhvani/downloader"
	"dhvani/filemanager"
	"dhvani/player"
	"fmt"
)

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
		player.Play(res[n-1])

	} else {

		fmt.Println("Downloading...")
		name, dest, err := downloader.Getfile(ytres[n-count-1][0])
		if err != nil {
			panic(err)
		}
		player.Play(filemanager.AddSong(name, dest))
		fmt.Println("Playing", name)

	}

}
