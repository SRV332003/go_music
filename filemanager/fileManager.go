package filemanager

import (
	"fmt"
	"log"

	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/SRV332003/go_music/models"

	"golang.org/x/term"
)

type Song = models.Song

func LenFiles() int {
	return len(files)
}

var files []Song
var mutex sync.Mutex

var userDir, _ = os.UserHomeDir()
var configFile = path.Join(userDir, ".config", "dhvani", "config.cnf")
var configDir = path.Join(userDir, ".config", "dhvani")
var MusicDir string

func init() {

	MusicDir = path.Join(userDir, "Music")
	config, err := os.OpenFile(configFile, os.O_RDONLY, 0666)

	if err != nil {

		os.MkdirAll(configDir, 0777)
		f, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.Println("Config file created.")

		dirname := path.Join(userDir, "Music")

		n, err := f.Write([]byte(dirname))

		if err != nil {

			os.Remove(configFile)
			panic(err)
		}
		log.Println(n, "bytes written to config file.")
		MusicDir = dirname

	} else {

		b := make([]byte, 200)
		n, err := config.Read(b)
		if err != nil {
			panic(err)
		}
		log.Println(n, "bytes read from config file. :", string(b))
		MusicDir = string(b[:n])

	}

	mutex = sync.Mutex{}

	err = filepath.Walk(MusicDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Println(MusicDir, "is not a valid directory.")
			panic(err)
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp3" {

			mutex.Lock()
			song := Song{}
			song.ID = len(files)
			song.Name = strings.ReplaceAll(info.Name(), "00 - ", "")
			song.Name = strings.TrimSpace(song.Name)
			song.Path = path

			fmt.Println(song)
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

func GetRandom() Song {
	if len(files) == 0 {
		return Song{}
	}
	song := GetSongByID(rand.Intn(len(files)))
	return song
}

func SetMusicDir(dir string) {

	if dir == MusicDir || dir == "" {
		return
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println(dir, "does not exist.")
		return
	}

	if dir == "default" {
		dir = path.Join(userDir, "Music")
	}

	f, err := os.OpenFile(configFile, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Truncate(0)

	_, err = f.Write([]byte(dir))

	if err != nil {
		panic(err)
	}
	MusicDir = dir
	refreshList()
	return
	// log.Println(n, "bytes written to config file.")
}

func refreshList() {
	mutex.Lock()
	defer mutex.Unlock()

	files = files[:0]

	err := filepath.Walk(MusicDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Println(MusicDir, "is not a valid directory.")
			panic(err)
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp3" {

			song := Song{}
			song.ID = len(files)
			song.Name = strings.ReplaceAll(info.Name(), "00 - ", "")
			song.Name = strings.TrimSpace(song.Name)
			song.Path = path

			fmt.Println(song)
			files = append(files, song)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

}
