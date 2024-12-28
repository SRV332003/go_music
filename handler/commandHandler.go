package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

func HandleCommand(command rune, key keyboard.Key) (success bool) {
	args := []string{}

	success = true

	switch key {
	case keyboard.KeyArrowRight:
		args = append(args, "10")
		skipHandler(args)
		success = false
	case keyboard.KeyArrowLeft:
		args = append(args, "-10")
		skipHandler(args)
		success = false
	case keyboard.KeyArrowUp:
		volumeUpHandler(args)
		success = false
	case keyboard.KeyArrowDown:
		volumeDownHandler(args)
		success = false
	case keyboard.KeySpace:
		pausePlayHandler(args)
		success = false
	}
	if !success {
		return
	}

	switch command {
	case '[':
		fmt.Printf("%c\n", command)
		speedDownHandler(args)
	case ']':
		fmt.Printf("%c\n", command)
		speedUpHandler(args)
	case 'c':
		fmt.Printf("%c\n", command)
		ClearScr()
	case 'p':
		fmt.Printf("%c\n", command)
		pausePlayHandler(args)
	case 's':
		fmt.Printf("%c\n", command)
		advSearchHandler(args)
	case 'r':
		fmt.Printf("%c\n", command)
		randomSongHandler(args)
	case 'q':
		fmt.Printf("%c\n", command)
		exitHandler(args)
	case 'l':
		fmt.Printf("%c\n", command)
		loopHandler(args)
	case 'n':
		fmt.Printf("%c\n", command)
		nextHandler(args)
	case 'h':
		fmt.Printf("%c\n", command)
		help()
	case '+':
		// fmt.Printf("%c\n", command)
		volumeUpHandler(args)
		success = false
	case '-':
		// fmt.Printf("%c\n", command)
		volumeDownHandler(args)
		success = false
	default:
		success = false
	}
	return
}

func HandleWordInputs() {
	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)
	str, err := bufioReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	str = str[:len(str)-1]
	str = strings.TrimSpace(str)
	str = strings.Split(str, "\\")[0]
	log.Println(str)

	HandleWordCommands(str)

}
