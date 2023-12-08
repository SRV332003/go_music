package main

import (
	"bufio"
	"dhvani/handler"
	"log"
	"os"
)

func main() {

	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)
	handler.ClearScr()
	for {

		os.Stdout.WriteString("dhvani> ")
		str, err := bufioReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		handler.HandleInput(str)
	}

}
