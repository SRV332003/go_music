package filemanager

import (
	// "fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type Song struct {
	ID   int
	Name string
	Path string
}

var files []Song

func Getfiles() {

}

func init() {
	log.Println("init Ran")

	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dirname := path.Join(userDir, "Music")

	mutex := sync.Mutex{}

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
