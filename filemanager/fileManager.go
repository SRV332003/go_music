package filemanager

import (
	// "fmt"
	"dhvani/player"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/term"
)

type Song struct {
	ID   int
	Name string
	Path string
}

func (s Song) String() string {
	return fmt.Sprintf("%3d. %30s...", s.ID+1, s.Name[:min(30, len(s.Name))])
}

func (song Song) Play() {

	_, err := player.Play(song.Path)
	if err != nil {
		log.Println(err)
	}
}

var files []Song
var mutex sync.Mutex

func init() {
	log.Println("init Ran")

	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dirname := path.Join(userDir, "Music")

	mutex = sync.Mutex{}

	err = filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp3" {

			mutex.Lock()
			song := Song{}
			song.ID = len(files)
			song.Name = strings.ReplaceAll(info.Name(), "00 - ", "")
			song.Name = strings.TrimSpace(song.Name)
			song.Path = path

			// fmt.Println(song)
			files = append(files, song)
			mutex.Unlock()
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

}

func ListFiles() {
	cols, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	n := cols / 40
	count := 0
	for _, song := range files {
		count += 1
		if count%n == 0 {
			fmt.Println(song)
		} else {
			fmt.Print(song, "\t")
		}
	}
	if count%n != 0 {
		fmt.Println()
	}
}

func GetSongByID(id int) Song {
	if id < len(files) && id > -1 {
		return files[id]
	}
	return files[0]
}

func AddSong(name string, path string) Song {
	var s Song
	mutex.Lock()
	s.ID = len(files)
	s.Name = name
	s.Path = path
	files = append(files, s)
	mutex.Unlock()
	return s
}
