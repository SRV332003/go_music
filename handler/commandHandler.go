package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func HandleCommand(command rune) (success bool) {
	args := []string{}

	success = true

	switch command {

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
	case 'e':
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

	HandleWordCommands(str)

}
